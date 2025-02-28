package main

import (
    "os"
    "time"
    "github.com/go-kit/kit/log"
    "github.com/go-kit/kit/log/level"
)

func main() {
    // Create a logger with timestamp and logfmt output
    logger := log.NewLogfmtLogger(os.Stdout)
    logger = log.With(logger, "ts", log.DefaultTimestampUTC)

    // Add log level filtering
    logger = level.NewFilter(logger, level.AllowInfo()) // Only allow "info" and above

    // Log messages with different levels
    logger.Log("level", "info", "msg", "Prometheus scrape completed", "duration", "2s")
    logger.Log("level", "debug", "msg", "Detailed scrape info") // This won't appear
}
