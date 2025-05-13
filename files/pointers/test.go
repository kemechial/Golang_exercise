package main

import "fmt"

// Function that modifies a value using a pointer
func increment(n *int) {
    *n++
}

// Struct with a pointer field
type Person struct {
    Name string
    Age  *int
}

func main() {
    // Declaring and assigning pointers
    var num int = 10
    var ptr *int = &num

    fmt.Println("Initial value:", num)
    fmt.Println("Pointer address:", ptr)
    fmt.Println("Pointer dereferenced value:", *ptr)

    // Modifying value via pointer
    *ptr = 20
    fmt.Println("Modified value:", num)

    // Passing pointer to function
    increment(&num)
    fmt.Println("Value after increment:", num)

    // Using pointers in structs
    age := 25
    person := Person{Name: "Kaan", Age: &age}
    fmt.Println("Person Name:", person.Name)
    fmt.Println("Person Age:", *person.Age)

    // Modifying struct field via pointer
    *person.Age = 30
    fmt.Println("Updated Person Age:", *person.Age)
}