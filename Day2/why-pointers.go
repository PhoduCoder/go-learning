package main

import (
	"fmt"
)

var age int = 20

func passbyValue(age int){
	fmt.Printf("Initial value when passed in function %v \n", age)
	age = age+10
	fmt.Printf("Later value in function %v \n", age)
}

func main() {

	fmt.Printf("Initial value when passed in main %v \n", age)
	passbyValue(age)
	fmt.Printf("Later value after function call in main %v \n", age)
}

//"With values such as int, bool, and string, when you pass them to a function,
// Go makes a copy of the value, and it’s the copy that’s used in the function. 
//This copying means that a change that’s made to the value in the function 
//doesn’t affect the value that you used when calling the function."