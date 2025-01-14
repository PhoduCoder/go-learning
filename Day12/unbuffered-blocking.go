package main

import "fmt"

func main() {
	ch := make(chan int) //Unbuffered channel since capacity not specified

	//Send data 
	go func(){
		fmt.Printf("Sending data\n")
		ch <- 48 //Sender is blocked
		fmt.Printf("Data Sent\n")
	}()

	fmt.Printf("Receiving data\n")
	num := <- ch //Receiver is unbl
	fmt.Printf("Received data %d\n", num)
}