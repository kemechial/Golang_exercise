package	main

import "fmt"

func main() {
	cars := []string{"Ford", "Chevy", "Toyota"} // slice of strings
	for index, car := range cars { // range returns the index and value
		fmt.Println(index, car) // print the index and value
	}
	
}