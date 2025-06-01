package main


import (
	"fmt"
)

func main() {
	// Define a slice of integers
	numbers := []int{1, 2, 3, 4, 5}

	// Calculate the sum of the integers in the slice
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	// Print the sum
	fmt.Println("The sum of the numbers is:", sum)

	greeting := "Hello, World!"

	println([]byte(greeting))
}