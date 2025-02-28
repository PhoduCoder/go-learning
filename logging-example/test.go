package main

import (
    "os"
    "github.com/go-kit/kit/log"
    "github.com/go-kit/kit/log/level"
)

func main() {
    logger := log.NewLogfmtLogger(os.Stdout)
    logger = log.With(logger, "ts", log.DefaultTimestampUTC)
    logger = level.NewFilter(logger, level.AllowError()) // Should allow info and above

    // Test all levels
    logger.Log("level", "debug", "msg", "This is debug")
    logger.Log("level", "info", "msg", "This is info")
    logger.Log("level", "warn", "msg", "This is warn")
    logger.Log("level", "error", "msg", "This is error")
}
