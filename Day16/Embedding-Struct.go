package main

import (
	"fmt"
)

// Created a general variable name of type string
type name string

// Create a struct variable called location
type location struct {
	x int
	y int
}

// Create a struct variable called size
type size struct {
	width  int
	height int
}

// Creating the EMBEDDED struct
type dot struct {
	name
	location
	size
}

//mbedding is different than having a field that is a struct type.
// When you embed, the fields from the embedded struct get promoted.
//Once promoted, a field acts as if it’s defined on the target struct.

func getDots() []dot {
	var dot1 dot

	dot2 := dot{}
	dot2.name = "A"
	dot2.x = 5
	dot2.y = 6
	dot2.width = 10
	dot2.height = 20

	dot3 := dot{
		name: "B",
		location: location{
			x: 13,
			y: 27,
		},
		size: size{
			width:  5,
			height: 7,
		},
	}

	dot4 := dot{}
	dot4.name = "C"
	dot4.location.x = 101
	dot4.location.y = 209
	dot4.size.width = 87
	dot4.size.height = 43

	//How the height is now embedded into the dot struct
	fmt.Println("Dot4 size height %d", dot4.height)
	// This means there are two ways to access with struct fields
	// one is by struct-name.type.fieldName or by struct-name.fieldName
	fmt.Println("Dot4 width is %d", dot4.size.width)

	return []dot{dot1, dot2, dot3, dot4}
}

func main() {
	dots := getDots()
	for i := 0; i < len(dots); i++ {
		fmt.Printf("dot%v: %#v\n", i+1, dots[i])
	}

}
