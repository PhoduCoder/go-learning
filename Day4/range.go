package main

import "fmt"

func main() {

    nums := []int{2, 3, 4}
	sum := 0
	
	//range on arrays and slices provides both the index and value for each entry.
	//Above we didn’t need the index, so we ignored it with the blank identifier _.
	
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

	//Sometimes we actually want the indexes though.
    for i, num := range nums {
        if num == 4 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

	//Range can also iterate just over the keys
    for k := range kvs {
        fmt.Println("key:", k)
    }

	//Range over string, iterates over Unicode code points
    for _, c := range "go" {
        fmt.Println(c)
    }
}