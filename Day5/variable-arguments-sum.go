//This is to create a variadic function 
// where one can create a sum of variable number of arguments

package main

import (
  "fmt"
)

func main() {
  i := []int{ 5, 10, 15}
  fmt.Println(sum(5, 4))
  fmt.Println(sum(1,2,3,4,5,6,7,8,9))
  fmt.Println(sum(i...))
  j := []int{0,2}
  fmt.Println(sum(j...))
}

func sum(args ...int) int{
	totalSum := 0
	for i:=0; i<len(args); i++ {
		totalSum += args[i]
	}
	return totalSum
	//fmt.Println("The total sum is %v\n", totalSum)
}