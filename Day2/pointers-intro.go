package main

import (
	"fmt"
)

func main(){

	// We will learn about three different ways of declaring pointers

	var point1 *int // This creates a pointer whose redirection value is nil

	name := "gaurav"
	point2 := &name // Another way , pointing to address of another variable

	point3 := new(float64) //Another way to create a pointer

	if point1 != nil {
		fmt.Printf("The value and type of point1 is %#v, %T ", point1, point1)
	} else{
		fmt.Printf("Pointer point1 was nil\n")
	}

	if point2 != nil {
		fmt.Printf("The value and type of point2 is %#v, %T \n", point2, point2)
	}

	if point3 != nil {
		fmt.Printf("The value and type of point3 is %#v, %T \n", point3, point3)
	}
	
}