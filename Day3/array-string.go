package main

import (
	"fmt"
)

func main() {

	words := [4]string{"I", "love", "you", "Krishna"}

	fmt.Printf("The words array is %v\n", words)

	// Safe to run the len method which is efficient in Go
	fmt.Printf("The length of the array of strings is %v \n", len(words))
	fmt.Printf("%v, %v, %v, %v\t\n", words[3], words[0], words[1], words[2])

}