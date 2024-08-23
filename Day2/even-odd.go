package main

import (
	"fmt"
)

func main(){

	choice := 101

	if (choice % 2 == 0){
		fmt.Printf("Number is even")
	} else{
		fmt.Printf("Number is odd")
	}
}