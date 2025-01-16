package main

import (
	"fmt"
	"log"
	"gopkg.in/yaml.v3"
)


//Convert YAML to struct - unmarshal
//Convert struct to YAML - marshal 

type Car struct {
	TopSpeed   int
	Name       string
	Cool       bool
	Passengers []string
	Hobbies []map[string]string
	PassengersAge map[string]int
	realHobbies map[string]Passengers
}

func main() {
	c := Car{
		TopSpeed:   150,
		Name:       "Porsche",
		Cool:       true,
		Passengers: []string{"gaurav", "shreya"},
		Hobbies: []map[string]string{{"Gaurav":"cricket", "Shreya":"painting"},{"abc":"def"}},
		PassengersAge: map[string]int{"Gaurav":34,"Shreya":31,"India":76},
		realHobbies: 
	}

	fmt.Printf("The initial struct is %v\n", c)
	fmt.Println("============*************==============")

	output,err := yaml.Marshal(c)
	//output is an []byte
	if (err!=nil){
		log.Fatal(err)
	}
	fmt.Println(output)
	fmt.Println("==============************=============")
	fmt.Println(string(output))
}