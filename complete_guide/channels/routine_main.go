package main

import (
	"fmt"
	"net/http"
	"time"
)


func main()  {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.twitter.com",
		"http://www.linkedin.com",
		"http://www.github.com",
		"http://www.reddit.com",
		"http://www.youtube.com",
		"http://www.instagram.com",
		"http://www.pinterest.com",		
	}


	c := make(chan string)

	for _, link := range links {
		go checkLink2(link, c)
	}

	//for i := 0; i < len(links); i++ {
	/*
	
	for {
		//fmt.Println(<-c)
		go checkLink2(<-c, c)
	}
		*/

	for l := range c {
		//time.Sleep(2 * time.Second) // Sleep for 2 seconds before checking the next link
		//go checkLink2(l, c)
		go func ()  {
			time.Sleep(2 * time.Second) // Sleep for 2 seconds before checking the next link
			checkLink2(l, c)
		}()
	}
}

func checkLink2(link string, c chan string) {

	_, err := http.Get(link)

	if err != nil {	
		fmt.Println(link, "is down!")
		c <- link
		return
	} else {
		println(link, "is up!")
		c <- link
	}
	
}


	