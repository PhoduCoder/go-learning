package main

import (
	"os"
	"fmt"
)

func main(){
	text := []byte ("We wanted to add another line to an existing file")

	err := os.WriteFile( "abc.txt", text, 0644) //This is a method to write without opening the file
	//For an existing file it does NOT append rather remove everything

	if err != nil{
		fmt.Println("Panic because of err\n", err)
	}	
}