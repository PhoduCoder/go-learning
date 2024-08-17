package main

import (
	"fmt"
	"time"
)

func main() {
	Debug, loglevel, replicas, runtime := true, "trace", 3, time.Now()

	fmt.Println(loglevel, replicas, runtime, Debug)
}
