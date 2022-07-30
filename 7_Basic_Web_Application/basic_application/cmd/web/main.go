package main

import (
	"basicWebApp/pkg/handlers"
	"fmt"
	"net/http"
)

// Port number of our website
const portNumber = ":8080"

// main is the entry point.
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting an application on port", portNumber[1:])
	http.ListenAndServe(portNumber, nil)
}
