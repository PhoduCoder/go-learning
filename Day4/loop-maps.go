package main

import (
	"fmt"
)

func main(){

	students := map[int]string{1:"Gaurav",2:"Shreya",3:"Student1"}

	for i,j := range (students){
		fmt.Printf("Id for %s is %v\n", j,i)
	}
}