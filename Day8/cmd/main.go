package main

//import "fmt"
import "Day8/pkg/shapes"

func main() {
	t := shapes.Triangle{Base: 1.5, Height: 3} 
	s := shapes.Square{Side: 4}
	c := shapes.Circle{Radius: 2}

	//Calling exported (capitalized) functions of a package
	// is by calling the package name.function name, followed by inputs that the 
	//function of package takes
	shapes.PrintShapeDetails(t,s,c)

}