package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order number %d!\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("Making pizza number %d, it will take %d seconds...\n", pizzaNumber, delay)

		//delay for a bit
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza %d!\n", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza %d!\n", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order number %d is ready!\n", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we're trying to make
	var i = 0

	// run for ever or until we receive a quit notification

	// try to make pizzas
	for {
		// try to make pizza
		currentPizza := makePizza(i)

		// decision
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			// we tried to make a pizza, we sent something to data channel
			case pizzaMaker.data <- *currentPizza:

			case quitChan := <-pizzaMaker.quit:
				// close channels
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func main() {

	// seed the random number generator
	// deprecated in Go 1.20
	// rand.Seed(time.Now().UnixNano())

	// print out a message
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order no. %d is out for delivery", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("*** The customer is really mad")
			}
		} else {
			color.Cyan("Done making pizzas")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing channel", err)
			}
		}
	}

	// print out the ending message
	color.Cyan("----------------")
	color.Cyan("Done for the day")

	color.Cyan("We made %d pizzas, but failed to make %d, with %d attempts in total", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		color.Red("It was an awful day...")
	case pizzasFailed >= 6:
		color.Red("It was not a very good day...")
	case pizzasFailed >= 4:
		color.Yellow("It was an okay day...")
	case pizzasFailed >= 2:
		color.Yellow("It was a pretty good day...")
	default:
		color.Green("It was a great day...")
	}
}
