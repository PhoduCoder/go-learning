package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	exampleFunction()
}

func exampleFunction() {
	fmt.Println("Before stack trace")
	debug.PrintStack() // Prints the stack trace
	fmt.Println("After stack trace")
}
