package main

func main() {

	cards := newDeck()
	cards.saveToFile("cards")
	cards.shuffle()
	cards.print()
}
