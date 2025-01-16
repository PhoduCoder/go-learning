package main

import (
	"fmt"
	"log"
	"gopkg.in/yaml.v3"
	"os"
)

type Passenger struct {
    Age     int      `yaml:passengers.age` //Important to annotate the keys with the corresponding yaml keys
    Hobbies string `yaml:passengers.hobbies`
}

type Car struct {
	TopSpeed   int      `yaml:"topspeed"`
	Name       string   `yaml:"name"`
	Cool       bool     `yaml:"cool"`
	Passengers map [string]Passenger `yaml:"passengers"`
}

func main() {
	// Load the file; returns []byte
	f, err := os.ReadFile("passengers-two.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Create an empty Car to be are target of unmarshalling
	var c Car

	// Unmarshal our input YAML file into empty Car (var c)
	if err := yaml.Unmarshal(f, &c); err != nil {
		log.Fatal(err)
	}

	// Print out the new struct
	fmt.Printf("%+v\n", c)
}