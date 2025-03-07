package main

import (
	"fmt"
)

func main() {
	given_string := "Gaurav"
	reverse_string := reverse(given_string)

	fmt.Printf("The reversed string is %s\n", reverse_string)
}

func reverse(str1 string) string {
	//byte_eq := []byte(str1)

	//fmt.Println(byte_eq) // prints [71 97 117 114 97 118]

	var revbyte []byte

	//var revStr string

	length := len(str1)
	fmt.Printf("Length is %d\n", length)

	for i := 0; i < length; i++ {
		revbyte[i] := str1[(length - 1 + i):(length - i)]
	}

	// last_char := str1[length-1 : length]
	// fmt.Printf(last_char)

	// second_last_char := str1[length-2 : length-1]
	// fmt.Printf(second_last_char)

	fmt.Printf(string(revbyte)) // prints Gaurav in reverse

	return "Reverse"
}
