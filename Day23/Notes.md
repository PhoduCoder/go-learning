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

##Notes about struct

Take a look below, this is a struct that takes the above Hash type and create a struct which denotes 
the ring in a consistent hash

type Ring struct {
    mu       sync.RWMutex  //for safe concurrent usage
    hash     Hash // to be able to use a variety of hash functions
    replicas int // Number of replicas that you want for a given node on the ring, a.k.a number of virtual nodes
    keys     []int // list of all hashed virtual node
    vnodeMap map[int]string // Map to represent virtual nodes back to physical node 
}


Now when initilaising the struct, we could use one of two ways - but we prefer the pointer based init

```
r := &Ring{
    replicas: 3,
    hash: crc32.ChecksumIEEE,
    keys: []int{},
    vnodeMap: map[int]string{},
}
```

r is actually a pointer to the struct 
We use this type of initialiation if we 
are creating struct which will be modified and/or copied around a lot
This makes it much more memory efficient 

