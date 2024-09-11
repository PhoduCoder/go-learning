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
	fmt.Printf("The ID of the person passed is %v and type %T \n", args, args)

	fmt.Printf("The name of the person passed is %v \n", getUser(args))
}

func getUser(id string) string{
	Users := map[string]string{
		"001":"Gaurav",
		"002":"Shreya",
		"003":"Gautam",
		"004":"Sonal",
		"005":"Janki",
		"006":"BkVerma",
		"007":"RitaVerma",
	}

	name, err := Users[id]
	if err == true {
		return name
	}else{
		return "Invalid User"
	}
}