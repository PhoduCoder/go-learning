package main

import (
	"fmt"
	"sync"
)

func main() {
	var to, from int
	to = 1
	from = 100

	fmt.Print("The values of to and fro are", %d, %d, to, from )

	var res *int //Declaring res as a pointer of int type

	wg := &sync.waitGroup() //wg is a pointer that stores the address for this
	//to reference this waitgroup, we will have to use *wg

	//Calling a Go coroutine sum
	go sum(to, from, wg, res)

	//Now we need to use wg.Add to increment coroutine counter
	wg.Add(1)

	wg.wait()

	fmt.Printf("Printing from Main coroutine\n")
	fmt.Printf("The value of result is %d", *res)

}

func sum(to int, from int, wg *sync.WaitGroup, res *int) {
	for i := from; i <= to; i++ {
		*res := *res + i
	}
	wg.Done()
}
