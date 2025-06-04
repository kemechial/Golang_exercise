package main

import (
	"fmt"
	"io"
	"os"
	"log"
	"net/http"
)
/*
func main() {
	response, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err) 
	}
	fmt.Printf("%s", body)
}

*/

func main() {
	response, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	
	bs := make([]byte, 99999)
	response.Body.Read(bs)
	fmt.Printf("%s", bs)
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Status:", response.Status)
	fmt.Println("Content Length:", response.ContentLength)
	fmt.Println("Header:", response.Header)
	fmt.Println("######### Copying response body to stdout #########")	
	// Copying response body to stdout
	io.Copy(os.Stdout, response.Body)
}