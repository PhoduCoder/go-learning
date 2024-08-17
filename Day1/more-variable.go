package main

import (
	"fmt"
)

var (
	// <name1> <type1> = <value1>
	debug bool = true
	age int = 34
	salary float32 = 15000
	name string = "Gaurav"
)

func main() {
	fmt.Println(debug, age, salary, name)
}