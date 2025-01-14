package main

//import "os/signal"

import "fmt"

func main(){

	//Initialize a channel that passes string
	ch := make(chan string, 1)

	go func (message string) {
		ch <- message //sending message to channel
		fmt.Printf("Is the buffer blocked\n")
	}("Hello from anonymous function")

	//Now we receive some information from the other side channel
	passed := <-ch
	fmt.Printf("No receiver configured\n")
	fmt.Printf("Passed is %s\n", passed)

}