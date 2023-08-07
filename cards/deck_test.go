package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades but got %v", d[0])
	}
	if d[len(d)-1] != "King of Clever" {
		t.Errorf("Expected last card to be King of Clever but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	d := newDeck()
	hand, _ := deal(d, 4)
	os.Remove("_testDeck")
	hand.saveToFile("_testDeck")
	loadedDeck := newDeckFromFile("_testDeck")

	if len(loadedDeck) != 4 {
		t.Errorf("Expected deck to have 4 card but got %v", len(loadedDeck))
	}
	os.Remove("_testDeck")
}
