//Organizing code more easily 

//Without receivers

type Car struct {
    speed int
}

func getSpeed(c Car) int {
    return c.speed
}

func main() {
    myCar := Car{speed: 100}
	fmt.Println(getSpeed(myCar))  // Outputs: 100
	//Note how we call the getSpeed method and pass car as an argument
	//See the difference in function calling via receivers
	//where it is more like the object.Function() much like OOPS
}
