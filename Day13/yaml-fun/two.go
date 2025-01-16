package main

import (
    "fmt"
    "log"
    "gopkg.in/yaml.v3"
)

type Passenger struct {
    Age     int
    Hobbies []string
}

type Car struct {
    TopSpeed int
    Name     string
    Cool     bool
    Passengers map[string]Passenger //key is string value is a struct
}

func main() {
    c := Car{
        TopSpeed: 150,
        Name:     "Porsche",
        Cool:     true,
        Passengers: map[string]Passenger{
            "Gaurav": {
                Age:     34,
                Hobbies: []string{"cricket", "chess"},
            },
            "Shreya": {
                Age:     31,
                Hobbies: []string{"painting"},
            },
            "Vijay": {
                Age:     45,
                Hobbies: []string{"Guitar"},
            },
        },
    }

    //fmt.Printf("The initial struct is %v\n", c)
    fmt.Println("---")

    output, err := yaml.Marshal(c)
    // output is an []byte
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(output))
}