package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/google/go-github/v62/github"
	"golang.org/x/oauth2"

	awsCfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	ecrTypes "github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

var (
	allowedEnvs = map[string]struct{}{"dev": {}, "uat": {}, "prod": {}}
	semverRE    = regexp.MustCompile(`^v?\d+\.\d+\.\d+(?:[-+].*)?$`)
)

func main() {
	var (
		owner        = flag.String("owner", "", "GitHub org/user that owns the repos")
		reposCSV     = flag.String("repos", "", "Comma-separated list of GitHub repository names (no owner)")
		mainBranch  = flag.String("main", "main", "Name of the main branch to branch from")
		awsRegion   = flag.String("region", os.Getenv("AWS_REGION"), "AWS region for ECR (defaults to $AWS_REGION)")
		registryID  = flag.String("registry", "", "ECR registry/account ID (optional; if empty will resolve via STS)")
		branchPref  = flag.String("branch-prefix", "auto/chart-bump", "Prefix for created branches")
		dryRun      = flag.Bool("dry", false, "Dry run: do not push commits/PRs")
	)
	flag.Parse()

	if *owner == "" || *reposCSV == "" {
		log.Fatalf("missing required flags: -owner and -repos")
	}
	if *awsRegion == "" {
		log.Fatalf("missing AWS region: set -region or AWS_REGION env var")
	}

	repos := splitCSV(*reposCSV)
	ctx := context.Background()

	// --- AWS clients ---
	cfg, err := awsCfg.LoadDefaultConfig(ctx, awsCfg.WithRegion(*awsRegion))
	if err != nil {
		log.Fatalf("load AWS config: %v", err)
	}
	acct := *registryID
	if acct == "" {
		acct, err = resolveAWSAccountID(ctx, cfg)
		if err != nil {
			log.Fatalf("resolve AWS account ID: %v", err)
		}
	}
	ecrClient := ecr.NewFromConfig(cfg)
	ecrHost := fmt.Sprintf("%s.dkr.ecr.%s.amazonaws.com", acct, *awsRegion)

	// --- GitHub client ---
	ghToken := os.Getenv("GITHUB_TOKEN")
	if ghToken == "" {
		log.Fatalf("GITHUB_TOKEN env var is required")
	}
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: ghToken})
	gh := github.NewClient(oauth2.NewClient(ctx, ts))

	branchName := fmt.Sprintf("%s/%s", strings.TrimSuffix(*branchPref, "/"), time.Now().UTC().Format("20060102-150405"))

	// Process repos sequentially; inside each repo we update files concurrently per addon/env if desired.
	for _, repo := range repos {
		if err := processRepo(ctx, gh, ecrClient, *owner, repo, *mainBranch, branchName, ecrHost, *dryRun); err != nil {
			log.Printf("repo %s: %v", repo, err)
		}
	}
}

func splitCSV(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func resolveAWSAccountID(ctx context.Context, cfg awsCfg.Config) (string, error) {
	client := sts.NewFromConfig(cfg)
	id, err := client.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return "", err
	}
	return *id.Account, nil
}

// processRepo discovers addon config.json files, computes desired chart versions from ECR, and opens a PR with updates.
func processRepo(ctx context.Context, gh *github.Client, ecrClient *ecr.Client, owner, repo, mainBranch, branchName, ecrHost string, dry bool) error {
	log.Printf("=== %s/%s ===", owner, repo)

	// 1) Get main ref SHA
	ref, _, err := gh.Git.GetRef(ctx, owner, repo, "refs/heads/"+mainBranch)
	if err != nil {
		return fmt.Errorf("get ref %s: %w", mainBranch, err)
	}
	mainSHA := ref.GetObject().GetSHA()

	// 2) Create working branch
	if !dry {
		_, _, err = gh.Git.CreateRef(ctx, owner, repo, &github.Reference{
			Ref: github.String("refs/heads/" + branchName),
			Object: &github.GitObject{SHA: github.String(mainSHA)},
		})
		if err != nil {
			// If branch exists, continue; else error
			if !strings.Contains(strings.ToLower(err.Error()), "reference already exists") {
				return fmt.Errorf("create branch %s: %w", branchName, err)
			}
		}
	}

	// 3) Walk tree to find addons/*/*/config.json
	tree, _, err := gh.Git.GetTree(ctx, owner, repo, mainSHA, true)
	if err != nil {
		return fmt.Errorf("get tree: %w", err)
	}

	type fileTarget struct {
		path   string
		addon  string
		env    string
		blobSHA string
	}
	var targets []fileTarget
	for _, e := range tree.Entries {
		p := e.GetPath()
		if !strings.HasPrefix(p, "addons/") || !strings.HasSuffix(p, "/config.json") { continue }
		parts := strings.Split(p, "/") // addons/<addon>/<env>/config.json
		if len(parts) != 4 { continue }
		addon := parts[1]
		env := parts[2]
		if _, ok := allowedEnvs[strings.ToLower(env)]; !ok { continue }
		targets = append(targets, fileTarget{path: p, addon: addon, env: strings.ToLower(env), blobSHA: e.GetSHA()})
	}
	if len(targets) == 0 {
		log.Printf("no addon config.json files found; skipping")
		return nil
	}

	// 4) For each target, compute desired chart version from ECR and update content
	chartRepoPrefix := "oci://" + ecrHost
	updatedCount := 0
	// Concurrency: process with a small worker pool to avoid API throttling
	sem := make(chan struct{}, 4)
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, target := range targets {
		wg.Add(1)
		sem <- struct{}{}
		go func(t fileTarget) {
			defer wg.Done()
			defer func() { <-sem }()

			version, digest, err := pickChartForEnv(ctx, ecrClient, t.addon, t.env)
			if err != nil {
				log.Printf("  %s: lookup ECR: %v", t.path, err)
				return
			}

			// Fetch current file to get SHA and content
			fc, _, _, err := gh.Repositories.GetContents(ctx, owner, repo, t.path, &github.RepositoryContentGetOptions{Ref: mainBranch})
			if err != nil { log.Printf("  %s: get contents: %v", t.path, err); return }
			oldContent, err := fc.GetContent()
			if err != nil { log.Printf("  %s: decode contents: %v", t.path, err); return }

			newJSON, changed, err := updateConfigJSON([]byte(oldContent), chartRepoPrefix, t.addon, version, digest)
			if err != nil { log.Printf("  %s: update JSON: %v", t.path, err); return }
			if !changed { log.Printf("  %s: already up-to-date", t.path); return }

			msg := fmt.Sprintf("chore(%s/%s): bump helm chart to %s (%s)", t.addon, t.env, version, shortDigest(digest))
			if dry {
				log.Printf("DRY-RUN would update %s => version=%s digest=%s", t.path, version, digest)
				mu.Lock(); updatedCount++; mu.Unlock()
				return
			}
			// Update via Contents API on our branch
			_, _, err = gh.Repositories.UpdateFile(ctx, owner, repo, t.path, &github.RepositoryContentFileOptions{
				Message: github.String(msg),
				Content: newJSON,
				Branch:  github.String(branchName),
				SHA:     fc.SHA,
				Author:  &github.CommitAuthor{Name: github.String("chart-bot"), Email: github.String("chart-bot@example.com")},
			})
			if err != nil { log.Printf("  %s: update file: %v", t.path, err); return }
			log.Printf("  updated %s -> %s (%s)", t.path, version, shortDigest(digest))
			mu.Lock(); updatedCount++; mu.Unlock()
		}(target)
	}
	wg.Wait()

	if updatedCount == 0 {
		log.Printf("no updates to propose; skipping PR")
		return nil
	}

	if dry { return nil }

	// 5) Open PR
	title := fmt.Sprintf("chore(addons): bump Helm charts (%d files)", updatedCount)
	body := "Automated bump: selects latest eligible Helm chart from ECR based on marker tags (dev: ready_to_release, uat: release_uat, prod: release_prod).\n\n" +
		"This PR was generated by a tool. Please verify values overlays as needed."
	pr, _, err := gh.PullRequests.Create(ctx, owner, repo, &github.NewPullRequest{
		Title: github.String(title),
		Head:  github.String(branchName),
		Base:  github.String(mainBranch),
		Body:  github.String(body),
	})
	if err != nil {
		return fmt.Errorf("create PR: %w", err)
	}
	log.Printf("opened PR #%d: %s", pr.GetNumber(), pr.GetHTMLURL())
	return nil
}

// pickChartForEnv chooses the newest semver tag overall for an addon (ECR repo == addon) and returns that tag and digest.
// Env gating is enforced separately by envAllowedForDigest.
// pickChartForEnv chooses the newest semver tag for an addon that is **eligible for the given env** based on marker tags.
// ECR repo name is assumed to equal addon.
func pickChartForEnv(ctx context.Context, ecrClient *ecr.Client, addon string, env string) (string, string, error) {
	repo := addon
	var nextToken *string
	bestV := (*semver.Version)(nil)
	bestDigest := ""
	marker := requiredMarker(env)

	for {
		out, err := ecrClient.DescribeImages(ctx, &ecr.DescribeImagesInput{
			RepositoryName: &repo,
			MaxResults:     int32Ptr(1000),
			NextToken:      nextToken,
		})
		if err != nil {
			var rnfe *ecrTypes.RepositoryNotFoundException
			if errors.As(err, &rnfe) {
				return "", "", fmt.Errorf("ECR repo %q not found", repo)
			}
			return "", "", err
		}
		for _, img := range out.ImageDetails {
			if len(img.ImageTags) == 0 || img.ImageDigest == nil { continue }
			// Check marker
			if !hasMarkerStrPtrs(img.ImageTags, marker) { continue }
			// Find highest semver tag on this digest
			var versions []*semver.Version
			for _, t := range img.ImageTags {
				if t == nil { continue }
				st := strings.TrimSpace(*t)
				if semverRE.MatchString(st) {
					if v, err := semver.NewVersion(strings.TrimPrefix(st, "v")); err == nil {
						versions = append(versions, v)
					}
				}
			}
			if len(versions) == 0 { continue }
			sort.Sort(semver.Collection(versions))
			v := versions[len(versions)-1]
			if bestV == nil || v.GreaterThan(bestV) {
				bestV = v
				bestDigest = *img.ImageDigest
			}
		}
		nextToken = out.NextToken
		if nextToken == nil { break }
	}
	if bestV == nil { return "", "", fmt.Errorf("no eligible semver-tagged images for env %q in ECR repo %q", env, repo) }
	return bestV.Original(), bestDigest, nil
}

func hasMarkerStrPtrs(tags []*string, marker string) bool {
	for _, t := range tags {
		if t == nil { continue }
		lt := strings.ToLower(*t)
		if lt == strings.ToLower(marker) || lt == strings.ToLower(marker)+"=true" {
			return true
		}
	}
	return false
}

// envAllowedForDigest checks if the digest has the required marker tag for the environment.
// We consider any of the following tags as markers:
//  - dev:  "ready_to_release" or "ready_to_release=true"
//  - uat:  "release_uat" or "release_uat=true"
//  - prod: "release_prod" or "release_prod=true"
// (envAllowedForDigest removed; eligibility is handled in pickChartForEnv)

func requiredMarker(env string) string {
	switch strings.ToLower(env) {
	case "dev":
		return "ready_to_release"
	case "uat":
		return "release_uat"
	case "prod":
		return "release_prod"
	default:
		return ""
	}
}

// updateConfigJSON updates known keys inside config.json.
// Expected keys (case-insensitive):
//  - helm_chart_repo
//  - helm_chart
//  - helm_chart_version
//  - helm_chart_sha
func updateConfigJSON(in []byte, repoPrefix, addon, version, digest string) ([]byte, bool, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(in, &m); err != nil {
		return nil, false, err
	}

	changed := false
	set := func(key string, val interface{}) {
		if cur, ok := m[key]; !ok || !equals(cur, val) {
			m[key] = val
			changed = true
		}
	}
	// Ensure keys exist and are updated
	set("helm_chart_repo", path.Join(repoPrefix, addon))
	set("helm_chart", addon)
	set("helm_chart_version", strings.TrimPrefix(version, "v"))
	set("helm_chart_sha", digest)

	if !changed { return in, false, nil }
	out, err := json.MarshalIndent(m, "", "  ")
	return out, changed, err
}

func equals(a, b interface{}) bool {
	switch av := a.(type) {
	case string:
		bv, ok := b.(string); if !ok { return false }
		return av == bv
	case float64:
		// JSON numbers come as float64; compare string->float if b is string-numeric
		bs, ok := b.(string); if ok {
			return fmt.Sprintf("%v", av) == bs
		}
		bf, ok := b.(float64); if ok { return av == bf }
	}
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

func shortDigest(d string) string {
	if strings.HasPrefix(d, "sha256:") && len(d) > 18 {
		return d[:18]
	}
	if len(d) > 12 { return d[:12] }
	return d
}

func int32Ptr(i int32) *int32 { return &i }
