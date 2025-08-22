package main

import (
	"fmt"
)

func hello_world() {
	fmt.Println("Printing hello world coroutine")
}

func main() {
	fmt.Println("Starting my Program")
	//time.Sleep(3 * time.Second)
	go hello_world()
	fmt.Println("Ending the program")
}

//The code starts by printing Starting my Program,
//then it calls the hello_world() function.
//Then, the execution goes straight to printing Ending the program
//without waiting for the hello_world() function to complete.
// No matter how long it takes to run the hello_world() function,
// the main() function will not care about the hello_world() function as
//these functions will run independently.

//So you will never see "Printing hello world coroutine run"
