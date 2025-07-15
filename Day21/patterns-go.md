Middleware Chaining Pattern

It lets you wrap core handler functions with extra behavior (like logging, metrics, tracing) without changing their code. It promotes:

Separation of concerns (business logic vs. instrumentation)

Reusability of wrappers across many handlers

Easy composition of multiple middleware layers

It’s a common and clean way to build extensible HTTP servers in Go.


```
func instrf(name string, next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Starting handler:", name)
        next(w, r)
        fmt.Println("Finished handler:", name)
    }
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello!")
}

http.HandleFunc("/hello", instrf("helloHandler", helloHandler))
```

Instead of 

```
func helloHandler(w http.ResponseWriter, r *http.Request) {
    // Before: instrumentation / logging inline
    fmt.Println("Starting handler: helloHandler")

    // Actual handler logic
    fmt.Fprintln(w, "Hello!")

    // After: instrumentation / logging inline
    fmt.Println("Finished handler: helloHandler")
}

http.HandleFunc("/hello", helloHandler)
```
