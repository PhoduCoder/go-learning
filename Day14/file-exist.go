package main

import (
	"fmt"
	"os"
)


func main(){

	file, err := os.Stat("abc.txt")

	if err != nil {
		fmt.Println("File doesn't exist")
	}

	fmt.Println(file.Size(), file.Mode(), file.Name(), file.IsDir(), file.ModTime())

	content, err_file := os.ReadFile("abc.txt") //Content is a byte slice 
	// Only do this for small files 
	// Entire contents are loaded at once

	if err_file != nil {
		fmt.Println("Panic error %s", err_file)
	}

	fmt.Println("The contents are %s", string(content))



}

