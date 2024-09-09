package main

import (
	"fmt"
)

//We will talk about Go slices and how it works internally
// Slices also have a hidden array to it 
//Size of slice is determined by len
//while size of hidden array is determined by cap
//When we used append function, depending on the capacity of the 
//underlying hidden array, Go in the background adds the new elements to 
// this undelying array if there is a space in hidden array.
// other wise Go copies the elements into a new hidden array and then
//points the slice to the new hidden array

//also note that there is a starting point for the slice
//which is the starting point of the array
//unless the slice was created by taking the last n elements of the array

func main(){
	array1 := [6]int{28,29,25,30,31,34}

	slice1 := array1[1:5]
	slice2 := array1[:5]
	fmt.Printf("The slice is %v and the length is %v and capacity is %v\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("The slice is %v and the length is %v and capacity is %v\n", slice2, len(slice2), cap(slice2))

	new_slice := append(slice1, 35, 36,37,38,39,40,41,42,43,44)
	fmt.Printf("The slice is %v and the length is %v and capacity is %v\n", new_slice, len(new_slice), cap(new_slice))

	//Now it is possible to control the hidden array 
//by using the make function
//in make function when creating a slice, make(<sliceType>, <length>, <capacity>)
//capacity is optional but used for the hidden array length


	new_sl := make([]int, 4, 6)
	fmt.Printf("The slice is %v of type %T and the length is %v and capacity is %v\n", new_sl, new_sl, len(new_sl), cap(new_sl))

	app_sl := append(new_sl,25,26)
	fmt.Printf("The slice is %v and the length is %v and capacity is %v\n", app_sl, len(app_sl), cap(app_sl))
	app_sl_2 := append(app_sl,27)
	fmt.Printf("The slice is %v and the length is %v and capacity is %v\n", app_sl_2, len(app_sl_2), cap(app_sl_2))

	new_test_sl := new_sl[1:4]
	fmt.Printf("The slice is %v and the length is %v and capacity is %v\n", new_test_sl, len(new_test_sl), cap(new_test_sl))

	slice_without_cap := make([]int, 4)
	fmt.Printf("The slice is %v of length %v and capacity %v \n", slice_without_cap, len(slice_without_cap), cap(slice_without_cap))
}





