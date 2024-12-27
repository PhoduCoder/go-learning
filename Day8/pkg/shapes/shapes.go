package shapes

import "fmt"

type Shape interface {
	area() float64
	name() string
}

type Triangle struct {
	Base float64
	Height float64
}

type Square struct {
	Side float64
}

type Circle struct {
	Radius float64
}

//Each of these methods implements the interface
//so it will have to implement the methods in the interface

func (t Triangle) area() float64 {
	var area float64
	area = (0.5 * t.Base * t.Height)
	return area
}

func (s Square) area() float64 {
	return(s.Side * s.Side)
}

func (c Circle) area() float64 {
	return (2.71*c.Radius*c.Radius)
}

func (t Triangle) name() string {
	return("Triangle")
}

func (s Square) name() string {
	return("Square")
}

func (c Circle) name() string {
	return("Circular")
}

//This function takes the shape interface as an input 
// and prints the area and name of that interface
func PrintShapeDetails(shapes ...Shape) {
	for _, item := range shapes {
		fmt.Printf("The area of %s is: %.2f\n", item.name(), item.area())
	}
}