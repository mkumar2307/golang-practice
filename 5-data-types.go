package main

import "fmt"

// Define a struct (custom data type)
type Person struct {
    Name string
    Age  int
}

func main() {
    // Integer types
    var a int = 42
    var b int8 = -100
    var u uint = 100

    // Floating point
    var pi float64 = 3.1415

    // Boolean
    var isGoFun bool = true

    // String
    var language string = "Go"

    // Array
    var primes [3]int = [3]int{2, 3, 5}

    // Slice
    fruits := []string{"apple", "banana", "cherry"}

    // Map
    scores := map[string]int{"Alice": 90, "Bob": 85}

    // Struct
    person := Person{Name: "John", Age: 30}

    // Pointer
    ptr := &a

    // Interface
    var anyType interface{} = "I can be anything"

    // Print values
    fmt.Println("Integer:", a, b, u)
    fmt.Println("Float:", pi)
    fmt.Println("Boolean:", isGoFun)
    fmt.Println("String:", language)
    fmt.Println("Array:", primes)
    fmt.Println("Slice:", fruits)
    fmt.Println("Map:", scores)
    fmt.Println("Struct:", person)
    fmt.Println("Pointer to a:", *ptr)
    fmt.Println("Interface:", anyType)
}