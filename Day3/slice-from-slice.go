package main

import (
	"fmt"
)

func main(){

	slice1 := []int{1:1, 3:5}
	new_slice := slice1[1:3] //Starts from index 1 to index (3-1)
	fmt.Printf("The slice is %v\n", slice1)
	fmt.Printf("The new slice is %v\n", new_slice)

	slice1[2]=10 // The new slice is just a view
	//Changing the original slice will also change the new slice

	fmt.Printf("The slice after assignment %v\n", slice1)
	fmt.Printf("The new slice after assignment %v\n", new_slice)


}