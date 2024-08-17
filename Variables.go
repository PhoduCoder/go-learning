package main

import "fmt"

//l := 4

var l int = 4 //Here we explicitly declare the type of variable as well
var m = 5 // Here the type of variable is implicitly infered

// The above implicit declaration won't work 
// since it is outside the function

func main() {
	var i, j int = 1, 2
	k := 3 // This type of declaration can only be done inside function
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
	fmt.Printf("The value of l is %d \n", l)
	fmt.Printf("The value of m is %d \n", m)
}
