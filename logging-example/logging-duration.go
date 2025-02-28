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

    // Simulate some work with time
    start := time.Now()
    time.Sleep(500 * time.Millisecond) // Simulate a 500ms task
    duration := time.Since(start)

    // Log messages with different levels
    logger.Log("level", "info", "msg", "Prometheus scrape completed", "duration", duration)
    logger.Log("level", "debug", "msg", "Detailed scrape info") // This won't appear
}
