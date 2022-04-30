package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck lenght of 52, but got %v.", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades as first card, but got %v.", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected King of Clubs as last card, but got %v.", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	f := "_decktesting"
	os.Remove(f)

	deck := newDeck()
	deck.saveToFile(f)

	loadedDeck := newDeckFromFile(f)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove(f)
}
