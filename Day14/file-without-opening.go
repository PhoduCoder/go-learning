package main

import (
	"os"
	"fmt"
)

func main(){
	text := []byte ("We wanted to add another line to an existing file")

	//func WriteFile(filename string, data []byte, perm os.FileMode) error
	//Note that the data passed is always a byte slice

	err := os.WriteFile( "abc.txt", text, 0644) //This is a method to write without opening the file
	//For an existing file it does NOT append rather remove everything

	if err != nil{
		fmt.Println("Panic because of err\n", err)
	}	

	new_text := []byte("This is another file being created")

	another_err := os.WriteFile("def.txt", new_text , 0644)

	if another_err != nil {
		fmt.Println("Panic because of err:%s \n", another_err)
	}
}