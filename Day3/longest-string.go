//Create a function that takes an slice of string
// and return the longest string 

package main

import (
	"fmt"
)

func longest(arr[]string) string {
	var longest string

	for i:=0; i<len(arr); i++{
		if len(arr[i]) > len(longest) {
			//fmt.Printf("Is %v greater than %v\n", arr[i], longest)
			longest=arr[i]
		}else{
			//fmt.Printf(" %v is NOT greater than %v\n", arr[i], longest)
		}
	}
	return longest
}

func main(){
	names := []string{"gaurav","revetment","usa","jersey","porsche","JPMorganChase","AmazonWebServices"}
	fmt.Printf("The longest name from the list %v is %v \t\n", names, longest(names))
}