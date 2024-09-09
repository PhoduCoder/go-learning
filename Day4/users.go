package main

import(
	"fmt"
)

func getUsers() map[string]string{

	users := map[string]string{
		"uid001": "userA",
		"uid002": "userB",
		"uid003": "userC",
		"uid004": "userD",
	}

	users["uid007"] = "SuperUser"

	//Checking if key exist in a map or not 
	superuser, key_exist := users["uid007"]
	
	if (key_exist == true){
		fmt.Printf("The user is %v user\n", superuser)
	} else {
		fmt.Printf("The user is not present\n")
	}

	return (users)
}

func main(){
	fmt.Println("The map is", getUsers())
}