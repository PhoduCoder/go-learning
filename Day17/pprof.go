package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Registers pprof handlers
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	// Start pprof server
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	fmt.Println("Starting the program...")

	// Simulate CPU-intensive task
	go computeFibonacci(40)

	// Simulate memory-intensive task
	go allocateMemory()

	// Simulate goroutine workload
	go launchWorkers(5)

	// Run for a while to allow profiling
	time.Sleep(60 * time.Second)
}

// computeFibonacci simulates CPU-intensive work
// Also recursion
func computeFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return computeFibonacci(n-1) + computeFibonacci(n-2)
}

// allocateMemory simulates high memory usage
func allocateMemory() {
	var data [][]byte
	for i := 0; i < 10; i++ {
		block := make([]byte, 10*1024*1024) // 10MB per block
		data = append(data, block)
		time.Sleep(time.Second)
	}
	fmt.Println("Memory allocation done")
}

// launchWorkers simulates goroutine load
func launchWorkers(count int) {
	for i := 0; i < count; i++ {
		go worker(i)
	}
}

// worker function keeps running indefinitely
func worker(id int) {
	for {
		fmt.Printf("Worker %d is running\n", id)
		time.Sleep(time.Second)
	}
}

// dumpHeapProfile captures heap profile at runtime
func dumpHeapProfile() {
	f, err := os.Create("heap_profile.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	runtime.GC() // Trigger GC before capturing heap profile
	pprof.WriteHeapProfile(f)
}
