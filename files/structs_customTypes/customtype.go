package main

import (
    "fmt"
    "strings"
    "math"
)

// Define a custom type based on string
type MyString string

// Attach a method to convert to uppercase
func (s MyString) ToUpper() string {
    return strings.ToUpper(string(s))
}

// Define a custom type based on float64
type MyFloat float64

// Attach a method to get the square root
func (f MyFloat) Sqrt() float64 {
    return math.Sqrt(float64(f))
}

func main() {
    // Demonstrate MyString type
    str := MyString("hello, kaan!")
    fmt.Println("Uppercase String:", str.ToUpper())

    // Demonstrate MyFloat type
    num := MyFloat(25.0)
    fmt.Println("Square Root:", num.Sqrt())
}