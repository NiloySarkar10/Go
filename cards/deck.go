package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Creates a new type of 'deck' which is a slice of strings with added functionality
type deck []string

// Creates and returns a new deck of playing cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Diamonds", "Spades", "Hearts", "Clubs"}
	cardValues := []string{"One", "Two", "Three", "Four", "Five"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Function available for all variables of type 'deck' which will loop over the slice
// and print the index and the value
func (d deck) printDeck() {
	for _, card := range d {
		fmt.Println(card)
	}
}

// Function to deal a hand when called. It returns two decks, first the hand dealt out
// and second the remaining cards in the deck.
func dealHand(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Function to take a deck and return a ',' separated string
func (d deck) deckToString() string {
	return strings.Join([]string(d), ",")
}

// Function to take in a filename and save the deck passed to system
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.deckToString()), 0666)
}
