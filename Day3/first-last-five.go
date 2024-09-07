package main

import (
	"fmt"
)

func main(){
	slice1 := []int{2,4,6,8,10,12,14,16,18,20}

	//Print first five
	fmt.Printf("The first five elements are %v, %v\n", slice1[0:5], slice1[:5])

	//Print last five
	fmt.Printf("The last five elements are %v\n", slice1[(len(slice1)-5):len(slice1)])
}