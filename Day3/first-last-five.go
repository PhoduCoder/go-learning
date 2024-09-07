package main

import (
	"fmt"
)

func main(){
	slice1 := []int{2,4,6,8,10,12,14,16}

	//Print first five
	fmt.Printf("The first five elements are %v, %v\n", slice1[0:5], slice1[:5])

	//  s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//  m += fmt.Sprintln("Last  :", s[len(s)-1], s[len(s)-1:len(s)], s[len(s)-1:])
	//Print last five
	fmt.Printf("The last five elements are %v\n", slice1[(len(slice1)-5):len(slice1)])
}