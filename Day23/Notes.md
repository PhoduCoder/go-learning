In Go, the  way to declare an empty struct is struct{}{}

##
In Go functions are first class value.
This means functions can be assigned to a variable
Functions can be passed as argument to other functions
You can also return a function from another function

```
package main
import "fmt"

func addOne(x int) int {
    return x + 1
}

func main() {
    f := addOne         // f is now a variable
    fmt.Println(f(5))   // 6
}
```

```
type Hash func([]byte) uint32 //declares a type Hash which represents any function that takes
// []bytes as input and outputs uint32

//So think of this type as with whom you can attach any function
// which has a similar signature ([]byte) uint32

//So this type accepts any function with this input output
//So Hash is like a label or alias for a whole class of functions with that shape.

func myHash(data []byte) uint32 {
    return uint32(len(data))
}

func main() {
    var h Hash = myHash           // assign function as a value
    fmt.Println(h([]byte("abc"))) // prints 3
}
```