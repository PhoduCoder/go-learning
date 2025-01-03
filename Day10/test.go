package main

import "fmt"

func rot13(s string) string {
    result := make([]byte, len(s))
    fmt.Printf("The initial result string is %s\n", result)
    for i := 0; i < len(s); i++ {
        char := s[i]
        fmt.Printf("The current processing character is %c\n", char)
        switch {
        case char >= 'a' && char <= 'z':
            fmt.Println("Inside the smaller case")
            result[i] = 'a' + (char-'a'+13)%26
        case char >= 'A' && char <= 'Z':
            fmt.Println("Inside the larger case")
            result[i] = 'A' + (char-'A'+13)%26
        default:
            result[i] = char
        }
    }
    return string(result)
}

func main() {
    var unicoded string
    unicoded = rot13("Gaurav")
    fmt.Printf("The unicoded value is %s\n", unicoded)
}
