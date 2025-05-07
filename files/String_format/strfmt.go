package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Basic string formatting using Sprintf
	name := "John"
	age := 25
	formattedStr := fmt.Sprintf("My name is %s and I am %d years old.", name, age)
	fmt.Println(formattedStr)

	// Using different formatting verbs
	num := 123.456
	fmt.Println(fmt.Sprintf("Default floating point: %f", num))
	fmt.Println(fmt.Sprintf("Scientific notation: %e", num))
	fmt.Println(fmt.Sprintf("Hexadecimal: %x", 255))

	// Padding and alignment
	fmt.Println(fmt.Sprintf("Right-aligned: |%10s|", name))
	fmt.Println(fmt.Sprintf("Left-aligned: |%-10s|", name))
	fmt.Println(fmt.Sprintf("Zero-padded number: %05d", age))

	// String concatenation
	concatStr := "Hello, " + "World!"
	fmt.Println(concatStr)

	// Converting numbers to strings
	number := 42
	strNumber := strconv.Itoa(number)
	fmt.Println("Converted number: " + strNumber)

	// Formatting booleans
	truth := true
	fmt.Println(fmt.Sprintf("Boolean value: %t", truth))
}