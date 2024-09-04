package main

import (
	"fmt"
)

func main(){
	fmt.Print(message()) //Note arguments can also be functions
}

func message() string {
	m := ""
	arr := [4]int{4,3,2,1}
	
	for i :=0; i < len(arr); i++ {
		arr[i]= arr[i] * arr[i]
	    m += fmt.Sprintf("%v: %v\n", i, arr[i])
	}
	return m
}
