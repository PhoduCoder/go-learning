//Your task is to print the user name based on the ID

//Assumptions
//The user name and Id are stored in a map called Users
//Our program passes the ID of the user as a CLI argument

//If Id doesn't exist, there must be an appropriate message
//else it should print the name of the person whose ID was passed.


//Calling the function is '''go run Challenge1.go 001'''

package main

import(
	"fmt"
	"os"
)

func main(){
	args := os.Args[1] //represent the ID of the person called
	//fmt.Printf("The ID of the person passed is %v and type %T \n", args, args)

	fmt.Printf("Hi %v \n", getUser(args))
}

func getUser(id string) string{
	Users := map[string]string{
		"001":"John",
		"002":"Mark",
		"003":"David",
		"004":"Justin",
		"005":"Tom",
		"006":"Chris",
		"007":"Michael",
		"008":"James",
		"009":"Steven",
		"010":"Brian",
		"011":"Adam",
		"012":"Joseph",
	}

	name, err := Users[id]
	if err == true {
		return name
	}else{
		return "Invalid User"
	}
}