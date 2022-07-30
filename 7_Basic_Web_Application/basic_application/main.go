package main

import (
	"fmt"
	"net/http"
)

// Port number of our website
const portNumber = ":8080"

// main is the entry point.
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Starting an application on port", portNumber[1:])
	http.ListenAndServe(portNumber, nil)
}
