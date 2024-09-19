//Define anonymous function is same as 
// defining a normal function 
// but without the name



package main

import(
	"fmt"
)

func main(){
	func(num int) {
		fmt.Printf("Coming from anonymous function\n")
		fmt.Printf("The squared number is %v\n", num*num)
	}(3) //Calling the anonymous function with a param
	fmt.Printf("This is coming from main function\n")
}