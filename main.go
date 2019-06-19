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
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	// Creating a channel instance, typed as string
	// Channel is a hub which allows main routine and child routine communication to each others.
	c := make(chan string)
	for _, link := range links {
		// Run this method on a Child Go Routine
		go checkLink(link, c)
	}

	// Loop contents in the Channel, infinitely, until the user press CTL+C	
	for l := range c {
		// Call the checklink function with 5 seconds declay by using an anonymous function
		// within child routine
		go func(link string) {
			// Add 5 seconds delay between method calls by using time.Sleep function
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		errMsg := link + " might be down."
		fmt.Println(errMsg)
		c <- link // Push link
		return
	}

	successMsg := link + " is up !"
	fmt.Println(successMsg)
	c <- link // Push link
}
