package main

import (
	"fmt"
	"encoding/json"
)

// Struct for Pod metadata
type Metadata struct {
	Name      string            `json:"name"` //This ensures during marshaling name appears and not Name as the field in Struct
	Namespace string            `json:"namespace"`
	Labels    map[string]string `json:"labels,omitempty"`
	annotations	string			//unexported fields that start with a small case can't be used outside the package
	// Only way to use unexported fields is by using  function receivers
	// but function receivers only work within the package
}

// Struct for Pod specification
type PodSpec struct {
	Containers []Container `json:"containers"`
}

// Struct for container
type Container struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

// Struct for Pod
type Pod struct {
	APIVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	Metadata   Metadata `json:"metadata"`
	Spec       PodSpec  `json:"spec"`
}

func main() {
	// Create a Pod object
	pod := Pod{
		APIVersion: "v1",
		Kind:       "Pod",
		Metadata: Metadata{
			Name:      "my-pod",
			Namespace: "default",
			Labels:    map[string]string{"app": "my-app"},
		},
		Spec: PodSpec{
			Containers: []Container{
				{Name: "nginx", Image: "nginx:latest"},
			},
		},
	}

	// Convert to JSON
	podJSON, err := json.MarshalIndent(pod, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(podJSON))
}
