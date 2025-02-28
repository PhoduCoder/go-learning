package main

import (
	"fmt"
)

func main(){

	//This will have operations with strings

	string1 := "Gaurav A"

	length := len(string1) //Gives us the length of any string

	fmt.Println(length)

	for i:=0; i<length; i++ {
		//fmt.Println(i)
		fmt.Println(string1[:i+1])
	}
}