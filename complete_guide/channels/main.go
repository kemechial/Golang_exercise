package main

import "net/http"

/*
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

	for _, link := range links {
		checkLink(link)
	}


}
*/
func checkLink(link string) {

	_, err := http.Get(link)

	if err != nil {	
		println(link, "is down!")
		return
	} else {
		println(link, "is up!")
	}
	
}