package main

import (
	"fmt"
	"net/http"
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

	// Print any messages that are pushed by child routines, into the channels 
	for i:= 0; i < len(links); i++ {
		// Pop message from the channel and then print it
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		errMsg := link + " might be down."
		c <- errMsg // Push error message to channel
		return
	}

	successMsg := link + " is up !"
	c <- successMsg // Push the success message into the channel
}
