package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" ---> Same as the next line
	card := newCard() // := needs to be used only for the first time declaration (indicate a new variable)
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
