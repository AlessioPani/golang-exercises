package main

import "fmt"

// Interface type
type bot interface {
	getGreeting() string
}

// Concrete types
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// We aren't using eb receiver, so we can omit the variable name and leave
// only the receiver type.
// func (eb englishBot) getGreeting() string {
func (englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	// Very custom logic for generating a spanish greeting
	return "Hola!"
}
