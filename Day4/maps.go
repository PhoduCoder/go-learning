package main

import(
	"fmt"
)

func main(){

	//Defining a map is like
	//map[key_value][value_type]{<key1>:<value1>, <key2>:<value2>}

	//Maps keys can only be int or string 
	//Maps values can be anything, even maps, slice or pointers

	age := map[string]int{"Gaurav":34,"Shreya":31,"India":76}

	//fmt.Printf("The age map is %v\n", age)
	fmt.Printf("Age of Gaurav is %v\n", age["G"])

}