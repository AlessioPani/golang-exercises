package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// The Dining Philosophers problem is well known in computer science circles,
// five philosophers, numbered from 0 through 4, live in a house where the
// table is laid for them; each philosopher has their own place at the table.
// Their only difficulty - besides those of philosophy - is that the dish
// served is a very difficult kind of spaghetti which has to be eaten with
// two forks. There are two forks next to each place, so that presents no
// difficulty. As a consequence, however, this means that no two neighbours
// may be eating simultaneosuly, since there are five philosophers and five forks.
//
// This is a simple implementation of Dijktra's solution to the
// "Dining Philosophers" dilemma.

// Philosopher is a struct which stores information about a philosopher
type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

// list of all philosophers
var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

// define global variables
var hunger = 3                  // how many times a person eat
var eatTime = 1 * time.Second   // how long it takes to eat
var thinkTime = 3 * time.Second // how long it takes to think
var mealMutex sync.Mutex        // a mutex for the mealOrder slice
var mealOrder []string          // the order in which philosophers finish their meal

func main() {
	// print out a welcome message
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("---------------------------")
	fmt.Println("The table is empty.")

	// start the meal
	dine()

	// print out finished message
	fmt.Println("The table is empty.")
	fmt.Printf("Order: %s", strings.Join(mealOrder, ", "))
}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all five forks
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start the meal
	for i := 0; i < len(philosophers); i++ {
		// fire off a go routine for the current philosopher
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// seat the philosopher to the table
	fmt.Printf("%s is seated at the table.\n", philosopher.name)
	seated.Done()

	seated.Wait()

	// eat 3 times
	for i := hunger; i > 0; i-- {
		// get a lock on both forks, starting from the lower fork to avoid
		// logical race condition
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork.\n", philosopher.name)
		}

		fmt.Printf("\t%s has both forks and is eating.\n", philosopher.name)
		time.Sleep(eatTime)
		fmt.Printf("\t%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()
		fmt.Printf("\t%s put down the forks.\n", philosopher.name)
	}

	fmt.Printf("%s is satisfied.\n", philosopher.name)
	fmt.Printf("%s left the table.\n", philosopher.name)
	mealMutex.Lock()
	mealOrder = append(mealOrder, philosopher.name)
	mealMutex.Unlock()
}
