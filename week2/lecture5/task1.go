package main

import "fmt"

type CardSuit = int 

const (
	club CardSuit = iota
	diamond
	heart
	spade
)
type Card struct {
	value int
	cardSuit CardSuit
}
func newCard(value int, suit CardSuit) Card {
	return Card{value, suit}
}
func compareCards(cardOne Card, cardTwo Card) int {
	if cardOne.value < 2 || cardOne.value > 13 || cardTwo.value < 2 || cardTwo.value > 13 {
		fmt.Println("Value is not in range [2,13]")
		return 3
	}
	if (cardOne.cardSuit > 4 || cardOne.cardSuit <0) || (cardTwo.cardSuit > 4 || cardTwo.cardSuit < 0) {
		fmt.Println("Suit is invalid.")
		return 3
	}
	if cardOne.value > cardTwo.value {
		return -1
	}
	if cardOne.value == cardTwo.value {
		if cardOne.cardSuit > cardTwo.cardSuit {
			return -1
		}
		if cardOne.cardSuit < cardTwo.cardSuit {
			return 1
		}
		return 0
	}
	if cardTwo.value > cardOne.value {
		return 1
	}
	return 3
}


func main() {
	// cardOne := newCard(8, 1)
	// cardTwo := newCard(7, 1)
	result := compareCards(newCard(20,1), newCard(7,1))
	if result == 0 {
		fmt.Println("The two cards are equal.")
	}
	if result == -1 {
		fmt.Println("The first card is greater than the second card.")
	}
	if result == 1 {
		fmt.Println("The second card is greater than the first card.")
	}
}