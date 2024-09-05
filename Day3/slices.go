package main

import (
	"fmt"
)

func main(){
	slice := []int{1:1,2:4}
	fmt.Printf("The slice is %v of type %T\n", slice, slice)

	//slice2 := []int{3,5}
	//Using append method 
	new_slice := append(slice, 5)

	fmt.Printf("The new slice is %v of type %T\n", new_slice, new_slice)
}