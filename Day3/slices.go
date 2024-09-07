package main

import (
	"fmt"
)

func main(){
	slice := []int{1:1,2:4}
	fmt.Printf("The slice is %v of type %T\n", slice, slice)

	//slice2 := []int{3,5}
	//Using append method 
	new_slice := append(slice, 5,)

	//You can add more than one value to append
	//You can even add slices to another slice using append method
	// locales = append(locales, extraLocales...)


	fmt.Printf("The new slice is %v of type %T\n", new_slice, new_slice)
}