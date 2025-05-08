package main

import (
	"fmt"
)

// Basic function
func greet() {
	fmt.Println("Hello, World!")
}

// Function with parameters and return value
func add(a int, b int) int {
	return a + b
}

// Function returning both integer quotient and remainder
func divideInt(dividend, divisor int) (quotient int, remainder int) {
	quotient = dividend / divisor
	remainder = dividend % divisor
	return
}

// Function returning float quotient and integer remainder
func divideFloat(dividend, divisor int) (quotient float64, remainder int) {
	quotient = float64(dividend) / float64(divisor) // Decimal division
	remainder = dividend % divisor                  // Integer remainder
	return
}

// Anonymous function assigned to a variable
var multiply = func(a, b int) int {
	return a * b
}

// Higher-order function (function that takes another function as a parameter)
func operate(a int, b int, op func(int, int) int) int {
	return op(a, b)
}

// Struct with a method
type Person struct {
	Name string
	Age  int
}

func (p Person) introduce() {
	fmt.Printf("Hi, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
	greet()

	sum := add(10, 5)
	fmt.Println("Sum:", sum)

	qInt, rInt := divideInt(10, 3)
	fmt.Printf("Integer Division -> Quotient: %d, Remainder: %d\n", qInt, rInt)

	qFloat, rFloat := divideFloat(10, 3)
	fmt.Printf("Float Division -> Quotient: %.2f, Remainder: %d\n", qFloat, rFloat)

	result := multiply(4, 5)
	fmt.Println("Multiplication:", result)

	opResult := operate(7, 3, add)
	fmt.Println("Operation result (addition):", opResult)

	person := Person{Name: "Alice", Age: 25}
	person.introduce()
}