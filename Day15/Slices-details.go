package main 

import (
	"fmt"
)

func AddOneToEachElement(slice []byte) {
    for i := range slice {
        slice[i]++
    }
}

var buffer [256]byte



func main() {
    slice := buffer[10:20] //Include 10th element, excludes the 20th element
    for i := 0; i < len(slice); i++ {
        slice[i] = byte(i+1)
    }
    fmt.Println("before", slice)
    AddOneToEachElement(slice)
    fmt.Println("after", slice)
}