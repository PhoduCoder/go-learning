//Define anonymous function is same as 
// defining a normal function 
// but without the name

package main

import(
	"fmt"
)

func main(){
	func(){ //Anonymous function is one without a name, hence anonymous
		fmt.Printf("Coming from anonymous function\n")
	}() //This is how you call the anonymous function
	fmt.Printf("This is coming from main function\n")
}