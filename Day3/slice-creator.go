package main

import (
	"fmt"
)

func main(){
	s1,s2,s3,s4 := slice_creator()
	fmt.Printf("s1: %v with len %v & capacity %v\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v with len %v & capacity %v\n", s2, len(s2), cap(s2))
	fmt.Printf("s3: %v with len %v & capacity %v\n", s3, len(s3), cap(s3))
	fmt.Printf("s4: %v with len %v & capacity %v\n", s4, len(s4), cap(s4))
}

func slice_creator() ([]int, []int, []int, []string) {
	var s1 []int

	//creating using make method
	s2 := make([]int, 4)

	//Specifying capacity of 6 as well
	s3 := make([]int, 4, 6)

	s4 := make([]string, 4, 6)

	return s1,s2,s3,s4
}