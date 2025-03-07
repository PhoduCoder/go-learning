package main

import (
	"fmt"
)

func main() {
	var str1 = "GauravInd"
	var str2 = "Shreya"

	result := alternateMerge(str1, str2)
	fmt.Println("Result is ", result)
}

func alternateMerge(str1 string, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)

	//var mergedString []string

	if len1 > len2 {
		for i := 0; i < len1; i++ {
			for j := 0; j < len2; j++ {
				//fmt.Println(str2[:j+1])
			}
			fmt.Println(str1[:i+1])
		}
	}

	return "Gaurav"
}
