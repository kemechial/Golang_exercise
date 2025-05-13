package main

import (
	"fmt"
)

// Person represents a basic struct with fields and methods.
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// FullName returns the person's full name.
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// IsAdult checks if the person is an adult.
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// Employee embeds the Person struct and adds additional fields.
type Employee struct {
	Person
	JobTitle string
	Salary   float64
}

// DisplayInfo prints details about the employee.
func (e Employee) DisplayInfo() {
	fmt.Printf("Name: %s\nAge: %d\nJob Title: %s\nSalary: %.2f\n",
		e.FullName(), e.Age, e.JobTitle, e.Salary)
}

// UpdateSalary modifies the salary of an employee using a pointer receiver.
func UpdateSalary(emp *Employee, newSalary float64) {
	emp.Salary = newSalary
}

func main() {
	// Creating an instance of Person
	p := Person{FirstName: "John", LastName: "Doe", Age: 30}
	fmt.Println("Full Name:", p.FullName())
	fmt.Println("Is Adult:", p.IsAdult())

	// Creating an instance of Employee
	e := Employee{
		Person:   Person{FirstName: "Alice", LastName: "Smith", Age: 25},
		JobTitle: "Software Engineer",
		Salary:   75000.00,
	}

	// Displaying employee information
	e.DisplayInfo()

	// Updating the salary using a function with a pointer parameter
	UpdateSalary(&e, 85000.00)

	// Displaying updated employee information
	fmt.Println("\nUpdated Employee Info:")
	e.DisplayInfo()
}