package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.com",
		"https://amazon.com",
	}

	c := make(chan string)

	// First iteration
	for _, link := range links {
		go checkLink(link, c)
	}

	// Continuous iteration
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second) // Wait for 5 sec
			checkLink(link, c)
		}(l)
	}
}

// Check the status of a link
func checkLink(l string, c chan string) {
	_, err := http.Get(l)

	if err != nil {
		fmt.Println(l, "may be down.")
		c <- l
		return
	}
	fmt.Println(l, "is up!")
	c <- l
}
