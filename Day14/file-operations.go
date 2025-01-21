package main

import (
	"os"
)

func main(){

	f, err := os.Create("abc.txt") //returns a file object
	if err != nil {
		panic(err)
	}

	//Write to file using either WriteString method and pass a string 
	// or by using Write method and pass a byte slice
	f.WriteString("Any string that we want to append\n")
	f.WriteString("Another line added\n") //Append string to the file 

	defer f.Close() //This closes the file at the end 

	f.Write([]byte ("This is the third line\n")) //converting a string to byte slice and then using the f.Write() method 


}