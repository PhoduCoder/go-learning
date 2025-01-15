package main

import (
	"fmt"
	"os"
	"os/signal"
	"flag"
	"syscall"
)

func main(){

	num1 := flag.Int("numOne", 100, "First number to be multiplied") 
	num2 := flag.Int("numTwo", 20, "Second number to be multiplied")

	flag.Parse()
	fmt.Printf("The two numbers you want multiplied are %d and %d\n", *num1, *num2)
	fmt.Printf("Please press CTRL+C to pass a signal like SIGINT\n")
	sig := make(chan os.Signal,1)
	done := make(chan struct{})
	//Catch the signal in your CLI
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL )

	go func(){
		for {
			s := <- sig // receiving the catched signal in s
			fmt.Printf("The processing has been interrupted with this signal %v\n", s)
			result := (*num1) * (*num2) //Note that num1 and num2 are both pointers
			fmt.Printf("The result of multiplying the passed two numbers is %d \n", result)
			done <- struct{}{} //sending empty struct to done channel
		}
	}()
	<- done //blocks the main and waits for the signal handler coroutine to signal it has handled the signal
	fmt.Println("Program is exiting")

}