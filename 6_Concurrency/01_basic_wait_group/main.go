package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {

	// WaitGroup to wait the end of each goroutines before exiting the program
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"eta",
		"theta",
		"zeta",
		"epsilon",
		"pi",
	}

	wg.Add(len(words))

	for i, w := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, w), &wg)
	}

	wg.Wait()

	// Add 1 to WaitGroup not to have a negative value
	// before calling wg.Done()
	wg.Add(1)
	printSomething("This is the second thing to be printed", &wg)
}
