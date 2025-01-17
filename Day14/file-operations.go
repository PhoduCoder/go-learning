package main

import (
	"os"
)

func main(){

	f, err := os.Create("abc.txt")
	if err != nil {
		panic(err)
	}

	f.WriteString("Any string that we want to append\n")
	f.WriteString("Another line added\n") //Append string to the file 

	defer f.Close() //This closes the file at the end 

	f.Write([]byte ("This is the third line\n")) //converting a string to byte slice and then using the f.Write() method 


}