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
	"strconv"
)

func main(){
	if len(os.Args) == 2 {
		args := os.Args[1] //represent the ID of the person called
		id, err := strconv.Atoi(args)
		if err == nil {
			//fmt.Printf("The ID of the person passed is %v and error %T \n", id, err)
			fmt.Printf("Hi %v \n", getUser(id))
		}
	} else {
		fmt.Printf("Call the program with right arguments , sample call '''go run Challenge1.go 001''' \n ")
	}
}

func getUser(id int) string{
	Users := map[int]string{
		1:"John",
		2:"Mark",
		3:"David",
		4:"Justin",
		5:"Tom",
		6:"Chris",
		7:"Michael",
		8:"James",
		9:"Steven",
		10:"Brian",
		11:"Adam",
		12:"Joseph",
	}

	name, err := Users[id]
	if err == true {
		return name
	}else{
		return "Invalid User"
	}
}