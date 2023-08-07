package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// create a new type of "deck" which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clever"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	hand := d[:handSize]
	remainingCards := d[handSize:]
	return hand, remainingCards
}

func (d deck) toString() string {
	stringSlice := []string(d)
	return strings.Join(stringSlice, ",")
}

func (d deck) saveToFile(fileName string) error {
	err := os.WriteFile(fileName, []byte(d.toString()), 0666)
	return err
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := string(bs)
	strSlice := strings.Split(s, ",")
	d := deck(strSlice)
	return d
}

func (d deck) shuffle() {
	deckLength := len(d)
	for i := range d {
		randIndex := rand.Intn(deckLength - 1)
		d[i], d[randIndex] = d[randIndex], d[i]
	}
}
