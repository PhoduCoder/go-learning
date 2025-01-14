package main

import "fmt"

func main(){
	s := make([]int, 10)
	//make takes a type to be created with length

	fmt.Println("The slice is %d", s)

	person := make(map[string]int, 5)
	person["Gaurav"] = 34
	person["Shreya"] = 31
 

	fmt.Printf("The map persons is %v actually\n", person)
}
