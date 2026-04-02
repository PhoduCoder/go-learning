package main

import "fmt'

func main(){

  //Declare a map 
  m := map[string]int{
    "apple":  5,
    "banana": 3,
}

m1 := make(map[string]int)

m1["apple"] = 5     // add
m1["banana"] = 3
m1["apple"] = 10    // update

value := m1["apple"]
fmt.Println(value)

fmt.Println(m["banana"])

//If the KEY doesn't exist, then the value comes as zero
fmt.Println(m["orange"])
  
}
