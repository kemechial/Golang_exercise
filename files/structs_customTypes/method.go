package main

import "fmt"

// Define a struct
type Person struct {
    Name string
    Age  int
}

// Method with a value receiver (does NOT modify the original struct)
func (p Person) ChangeName(newName string) {
    p.Name = newName
    fmt.Println("Inside ChangeName (value receiver):", p.Name)
}

// Method with a pointer receiver (modifies the original struct)
func (p *Person) ChangeNamePointer(newName string) {
    p.Name = newName
}

// Method with a pointer receiver (modifies the struct)
func (p *Person) HaveBirthday() {
    p.Age++
}

func (p Person) PrintName() {
	fmt.Println("Name:", p.Name)
}

func main() {
    // Creating an instance of Person
    p := Person{Name: "Kaan", Age: 25}

    p.PrintName()

    // Using the value receiver method (DOES NOT change Name)
    p.ChangeName("John")
    fmt.Println("After ChangeName method (value receiver):", p.Name) // Still "Kaan"

    // Using the pointer receiver method (CHANGES Name)
    p.ChangeNamePointer("John")
    fmt.Println("After ChangeNamePointer method (pointer receiver):", p.Name) // Now "John"

    fmt.Println("\nOriginal Age:", p.Age)

    // Using the pointer receiver method to modify the Age
    p.HaveBirthday()
    fmt.Println("After HaveBirthday method:", p.Age) // Age increases by 1
}