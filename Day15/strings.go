package main

import (
	"fmt"
)

func main() {

	//This will have operations with strings

	string1 := "Gaurav A"

	length := len(string1) //Gives us the length of any string

	fmt.Printf("The length of the string1 is %d \n", length)

	fmt.Println("Now printing the characters of the string one by one")

	for i := 0; i < length; i++ {
		//fmt.Println(i)
		fmt.Println(string1[:i+1])
	}
}
