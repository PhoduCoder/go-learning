package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x)) 
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	var n=sqrt(2)
	var m=sqrt(-4)
	fmt.Printf("N and m are %T, %v, %T, %#v", n, n, m, m)
}
