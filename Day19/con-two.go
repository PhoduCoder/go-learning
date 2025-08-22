package main

import "fmt"

func main() {
	var from int = 1
	var to int = 100
	res := calculate_sum(from, to)

	fmt.Printf("The result is %d ", res)
}

func calculate_sum(from, to int) int {
	result := 0
	for i := from; i <= to; i++ {
		result = (result + i)
	}
	return result
}

// Function that sums two number without any concurrency
