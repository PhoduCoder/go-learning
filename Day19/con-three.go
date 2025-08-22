package main

import (
	"fmt"
	"time"
)

func main() {
	var from int = 1
	var to int = 10000000000

	var res int
	go func() {
		res = calculate_sum(from, to)
	}() //Anonymous function coroutine since they have no names
	time.Sleep(5 * time.Second) //Makes the code sleep for 5 sec

	var res2 int
	res2 = calculate_sum(1, 100)

	fmt.Printf("The result is %d\n ", res)
	fmt.Printf("The result is %d \n", res2)
}

func calculate_sum(from, to int) int {
	result := 0
	for i := from; i <= to; i++ {
		result = (result + i)
	}
	return result
}

// Function that sums two number with  concurrency
//When time.Sleep is added
// (base) MacBookPro:Day19 alieninvader$ go run con-three.go
// The result is 55
//  The result is 5050

//WHEN TIME.SLEEP is not added
// (base) MacBookPro:Day19 alieninvader$ go run con-three.go
// The result is 0
//  The result is 5050
