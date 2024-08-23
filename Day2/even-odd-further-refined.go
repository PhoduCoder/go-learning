package main

import (
	"fmt"
	"errors"
)

func validate(input int) error{

	//error is  a built-in Go type used for errors
    //return a new error using the New function in the errors package
	if (input == 0){
		return errors.New("Zero is left aside, theoretically even")
	} else if (input < 0){
		return errors.New("Even odd debate doesn't apply to negative numbers")
	} else {
		return nil
	}
}

func main(){
	input := -10

    // check that the error is not equal to nil 
	// Here instead of a simple boolean, we have an assignment and a condition in IF
	if err := validate(input); err!= nil{
		fmt.Println(err)
	}else if (input % 2 == 0){
		fmt.Printf( " %#v is Even number", input )
	}else{
		fmt.Printf(" %#v is Odd number", input)
	}
}