package main

import (
    "fmt"
    "io"
    "log"
    "os"
)
func main() {
    f, err := os.Open("abc.txt")
    if err != nil {
        log.Fatalf("unable to read file: %v", err)
    }
    buf := make([]byte, 1)
    for {
        n, err := f.Read(buf)
        if err == io.EOF {
            break
        }
        if err != nil {
            fmt.Println(err)
            continue
		}
		// fmt.Println(string(buf[:n]))
    if n > 0 {
			//fmt.Println(string(buf[:n]))
			fmt.Print(string(buf[:n])) // By default, Go's fmt.Print buffers the output when writing to the terminal or console, especially for small chunks of data like a single character.


    }
    }
}