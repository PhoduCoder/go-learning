//This is an example of variadic function
// with multiple parameters 
//Note that variadic param must come last 

//In variadic functions func f(args ...Type) args is a slice of Type
//Hence you can do things like cap and len

package main

import (
	"fmt"
)

func main(){
	nums(2,20)
	nums(5,20,100)
	nums(3)
	//nums(4, []int{1:1,2:4})
	//The variadic params are converted to a slice,
	//one can't however pass a slice
	
	//If one wanted to use, one has to use the unpack operator
	slice1 := []int{1:1,2:4}
	nums(3, slice1...)
}

func nums(rank int, num ...int){
	fmt.Println("The rank is %v and num is %v", rank, num)
	fmt.Println("The length and capacity for the num slice is", len(num), cap(num))
}