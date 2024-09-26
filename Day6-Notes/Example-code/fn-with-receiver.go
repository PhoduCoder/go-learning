type Car struct {
    speed int
}

//This is a pointer receiver
func (c *Car) GetSpeed() int {
    return c.speed
}

func main() {
    myCar := Car{speed: 100}
    fmt.Println(myCar.GetSpeed())  // Outputs: 100
}