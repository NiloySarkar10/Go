package main

import "fmt"

func main() {
	cards := newDeck()
	cardsInHand, cardsRemainingInDeck := dealHand(cards, 5)
	errorFromInHandsCardSave := cardsInHand.saveToFile("in_hand_cards")
	errorFromInDeckCardSave := cardsRemainingInDeck.saveToFile(("cards_in_deck"))
	fmt.Println("errorFromInHandsCardSave: ", errorFromInHandsCardSave)
	fmt.Println("errorFromInDeckCardSave: ", errorFromInDeckCardSave)
	cards.saveToFile("my-cards")
	newDeck := newDeckFromFile("my-cards")
	fmt.Println("New deck: ", newDeck)
}
