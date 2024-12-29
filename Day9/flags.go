package main

import (
	"fmt"
	"flag"
	//"os"
)

func main(){
	//arguments := os.Args
	//fmt.Printf("The arguments passed are %s \n", arguments[1])

	//Note that the returned type is a string pointer
	wordPtr := flag.String("str", "Ram", "First argument which represent name of god" )

	numPtr := flag.Int("rep", 108, "Number of times to repeat the name")
	flag.Parse()

	fmt.Printf("My favorite god is %s\n", *wordPtr)

	fmt.Printf("%s name should be repeated %d times \n", *wordPtr, *numPtr)

	//Running this is by executing the below steps
	// go build flags.go
	// ./flags --str=Shiva

	// ./flags -h => will print all the available options

}
