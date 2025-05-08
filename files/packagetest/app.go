package main

import (
  "fmt"
	"example.com/packagetest/calculator"
)

func main() {
  maintest()
  var a, b int
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)

	fmt.Println("Addition:", calculator.Add(a, b))
	fmt.Println("Absolute Difference:", calculator.Subtract(a, b))
}
