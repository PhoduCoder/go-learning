package main

import "fmt"

type Person struct {
	Name string
	Age int
	Gender string
}

func main(){

	p1 := new(Person)
	//new is a built in function 
	//that allocates memory for 
	//zeroed values of a given type

	p1.Age = 34

	fmt.Printf("The new person Age is %d \n", p1.Age)


}