package main

import (
	"fmt"
	"io" // For io.ReadAll
	"os"
)

func main() {
	// Define the file name
	fileName := "example.txt"

	// Write to the file
	message := "Hello, Go file operations!"
	err := os.WriteFile(fileName, []byte(message), 0644) // Replaced ioutil.WriteFile
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File written successfully!")

	// Read from the file
	file, err := os.Open(fileName) // Open the file to get an io.Reader
	if err != nil {
		fmt.Println("Error opening file for reading:", err)
		return
	}
	defer file.Close() // Ensure the file is closed

	data, err := io.ReadAll(file) // Use io.ReadAll to read from the io.Reader
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	fmt.Println("File content:", string(data))
}