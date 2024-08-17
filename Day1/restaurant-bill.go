package main

import (
	"fmt"
	//"time"
)

func main() {
	var totalBill float64 = 0

	// Food cost for two items worth 13 USD
	totalBill = totalBill + (2 * 13)
	// Note that if we did something like below
	// that won't work since it will redeclare and reinitialize the variable
	//totalBill := totalBill + (2*13)


	fmt.Printf("Bill after Food %v \n", totalBill)

	//Drinks cost for 4 drinks worth 4 USD
	totalBill = totalBill + (4 * 4)
	fmt.Printf("Bill after food and drinks, %v \n", totalBill)

	//Customer is getting a 3 USD discount
	totalBill = totalBill -  3
	fmt.Printf("Bill after discount is %#v \n", totalBill)

	//Leave a 10 percent tip
	tip := 0.1 * (totalBill)

	//Apply 6.5% taxes for state and federal
	tax := 0.065*2*(totalBill)
	totalBill = (totalBill + tip + tax)
	fmt.Printf("The final bill after tip and tax is %#v and %T \n", totalBill, totalBill)

}