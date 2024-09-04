package main

import (
	"fmt"
)

func main() {

	//Initialize the array as
	// [<size>]<type>{<value1>,<value2>,…<valueN>}
	var arr1 [5]int
	fmt.Printf("The first array is %v\n", arr1) //Default values are default values of corresponding types

	//other ways to declare or initialize an array
	arr2 := [5]int{1, 3:5}

	fmt.Printf("The second array is %v\n", arr2)

	arr3 := [...]int{1,0,0,5,0}
	fmt.Printf("The third array  is %v\n", arr3)

	if (arr2 == arr3){
		fmt.Printf("Yeah comparison in array is simple and arr2 and arr3 are same is %v\n", (arr2 == arr3))
	}

}