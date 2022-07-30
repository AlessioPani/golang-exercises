package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the homepage handler.
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is my homepage!")
}

// About is the about page handler.
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	fmt.Fprintf(w, "This is my About page and 2 + 2 is %d.", sum)
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100, 0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0.")
		return
	}
	fmt.Fprintf(w, "%f divided by %f is %f.", float32(100), float32(0), f)
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by 0")
		return 0, err
	}
	result := x / y
	return result, nil
}

// addValues adds two integers and returns the sum.
func addValues(x, y int) int {
	return x + y
}

// main is the entry point.
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println("Starting an application on port", portNumber)
	http.ListenAndServe(portNumber, nil)
}
