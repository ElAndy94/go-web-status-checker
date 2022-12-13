package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://amazon.com",
	}

	c := make(chan string) // Create a channel

	for _, link := range links {
		go checkLink(link, c) // go routine trigger
	}

	for l := range c {
		go func(link string) { // Go Annonymous function
			time.Sleep(5 * time.Second) // Sleep for 5 seconds
			checkLink(link, c)
		}(l) // Go routine trigger
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link // Send the link to the channel
		return
	}

	fmt.Println(link, "is up!")
	c <- link // Send the link to the channel
}
