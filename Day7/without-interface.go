//Illustrating the concept of polymorhism
//This example since it doesn't define an interface 
//leads to repetition

package main

import "fmt"

type Dog struct{}
func (d Dog) Speak() string { return "Woof!" }

type Cat struct{}
func (c Cat) Speak() string { return "Meow!" }

func AnnounceDog(d Dog) {
	fmt.Println(d.Speak())
}

func AnnounceCat(c Cat) {
	fmt.Println(c.Speak())
}

func main() {
	d := Dog{}
	c := Cat{}

	AnnounceDog(d)
	AnnounceCat(c)
}
