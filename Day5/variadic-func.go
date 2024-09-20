package main

import (
	"fmt"
)

//Variadic function takes a variable 
// number of arguments

//Good to use it when the number 
// of arguments of a specified type 
// is unknown


//func f(parameterName …Type)

// The three dots in the front is called the pack operator
// Tells Go to store all the arguments of "Type" in parameterName

func main(){
	square(10)
	square(10,20,30)
	square()
}

func square(num ...int){
	fmt.Println(num)
}

// Output is [10]
// [10 20 30]