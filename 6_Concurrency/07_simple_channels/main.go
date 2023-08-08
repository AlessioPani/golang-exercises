package main

import (
	"fmt"
	"strings"
)

// shout has two parameters: a receive only chan ping, and a send only chan pong.
// Note the use of <- in function signature. It simply takes whatever
// string it gets from the ping channel,  converts it to uppercase and
// appends a few exclamation marks, and then sends the transformed text to the pong channel.
func shout(ping <-chan string, pong chan<- string) {
	for {
		s, ok := <-ping
		if !ok {
			// do something
			// ok is a boolean parameter. True if a returned value is actually
			// a value sent to the channel; False if the value comes from a
			// closed channel.
		}
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	// create two channels
	ping := make(chan string) // RX
	pong := make(chan string) // TX

	// start a goroutine
	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to quit)")

	for {
		// print a prompt
		fmt.Print("-> ")

		// get user input
		var userInput string
		fmt.Scanln(&userInput)

		if strings.ToLower(userInput) == "q" {
			break
		}

		// send userInput to "ping" channel
		ping <- userInput

		// wait for a response from the pong channel. Again, program
		// blocks (pauses) until it receives something from
		// that channel.
		response := <-pong

		// print the response to the console.
		fmt.Println("Response:", response)
	}

	fmt.Println("All done, closing channels...")

	// close the channels
	close(ping)
	close(pong)
}
