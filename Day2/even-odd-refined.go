package main

import (
	"fmt"
)

func main(){
	choice := -01

	if (choice == 0){
		fmt.Printf("Number is zero")
	} else if (choice < 0 ){
		fmt.Printf("Number can't be non negative")
	} else if (choice % 2 == 0){
		fmt.Printf(" number Is even")
	} else {
		fmt.Printf("Number is odd")
	}
}