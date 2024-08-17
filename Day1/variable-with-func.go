package main

import (
	"fmt"
	"time"
)

func getContents() (bool, string, time.Time, int) {

	return true, "gaurav", time.Now(), 34
}

func main() {
	isHuman, name, runTime, age := getContents()

	fmt.Println(isHuman, name, runTime, age)
}