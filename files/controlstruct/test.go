package main

import (
	"fmt"
)

func main() {
	// Conditional Statements (if, else if, else)
	num := 10
	if num < 5 {
		fmt.Println("Number is less than 5")
	} else if num > 5 && num < 15 {
		fmt.Println("Number is between 5 and 15")
	} else {
		fmt.Println("Number is 15 or more")
	}

	// Loop (for with condition)
	fmt.Println("Counting with a for loop:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Loop (for range)
	fruits := []string{"Apple", "Banana", "Cherry"}
	fmt.Println("Iterating over a slice:")
	for index, fruit := range fruits {
		fmt.Printf("Index %d: %s\n", index, fruit)
	}

	// Switch case
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("Start of the workweek!")
	case "Friday":
		fmt.Println("Weekend is near!")
	default:
		fmt.Println("Just another day.")
	}

	// Infinite Loop (break example)
	counter := 0
	for {
		fmt.Println("Looping...")
		counter++
		if counter == 3 {
			break
		}
	}

	// Continue example
	fmt.Println("Even numbers only:")
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}