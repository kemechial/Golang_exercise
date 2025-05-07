package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	var url string
	fmt.Print("Enter a website address: ")
	fmt.Scanln(&url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code", resp.StatusCode)
		return
	}

	var bodyBuilder strings.Builder
	_, err = io.Copy(&bodyBuilder, resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("First 10 lines of HTML content:")
	scanner := bufio.NewScanner(strings.NewReader(bodyBuilder.String()))
	for i := 0; i < 10 && scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning lines:", err)
	}
}

