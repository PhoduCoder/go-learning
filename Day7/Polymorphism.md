#Polymorphism in Go

Polymorphism is defined as the ability to appear in different forms. In Go there is no concept of classes, and one would implement polymorphism using interfaces. 

So when we define an interface, we just list the function signature of the methods that the interface provide.
Now any data type that implements that method is said to be implementing that interface.

Say we want to decribe a function that can function in three different ways depending on the type passed to it.
So say we have a speak function and to different types we could pass the speak function and we would get a different output.

So Cat will speak meow
Dog will speak bark
and a Human will speak words

Now if we used interface instead and created a function that accepts interface, one can use the concept of polymorphism by only defining the function once and using it differently with different 
passed types

=========

Expanding on  the above concept of creating functions that accept interface, one can also create structs that have interface type, it keeps the struct more flexible and reusable

========
```
type Storage interface {
	Save(data string)
}

//Two types that implements interface
type FileStorage struct{}
func (f FileStorage) Save(data string) { fmt.Println("Saving to file:", data) }

type MemoryStorage struct{}
func (m MemoryStorage) Save(data string) { fmt.Println("Saving to memory:", data) }

//Now a struct that uses that interface
type Service struct {
  name string
	storage Storage  // using interface
}

//This allows
s1 := Service{storage: FileStorage{}}
s2 := Service{storage: MemoryStorage{}}


```
=================
Instead if we used Service struct with 

```
type Service struct {
	storage FileStorage  // fixed to FileStorage
}

//you can’t use MemoryStorage without changing the struct.
```
