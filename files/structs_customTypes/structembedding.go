package main

import "fmt"

// Defining the Person struct
type Person struct {
    Name string
    Age  int
}

// Constructor function for Person
func NewPerson(name string, age int) Person {
    return Person{Name: name, Age: age}
}

// Method for Person struct
func (p Person) Greet() {
    fmt.Println("Hello, my name is", p.Name)
}

// Employee struct embeds Person
type Employee struct {
    Person
    Position string
}

// Constructor function for Employee
func NewEmployee(name string, age int, position string) Employee {
    return Employee{
        Person:  NewPerson(name, age), // Calling Person constructor
        Position: position,
    }
}

func main() {
    // Creating an Employee instance using the constructor
    emp := NewEmployee("Kaan", 30, "Software Engineer")

    // Accessing fields
    fmt.Println("Employee Name:", emp.Name) // Inherited from Person
    fmt.Println("Employee Age:", emp.Age)
    fmt.Println("Employee Position:", emp.Position)

    // Calling inherited method
    emp.Greet()
}