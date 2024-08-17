package main

import (
	"fmt"
)

func main() {
	count := 10

	fmt.Printf("Initial Count was %v \n", count)

	//Reduce the number by 1
	count --
	fmt.Printf("Now Count is %v \n", count)

	//Add 5 to count 
	count += 5
	fmt.Printf("Now Count is %v \n", count)

	//Remove 2 from count
	count -= 2
	fmt.Printf("Now Count is %v \n", count)

	//Add 1 from count
	count++
	fmt.Printf("Now Count is %v \n", count)

	//Add 10 to count 
	count -= -10
	fmt.Printf("Now Count is %v \n", count)

}