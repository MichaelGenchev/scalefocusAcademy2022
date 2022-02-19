package main

import "fmt"

type CardSuit = int 

const (
	club CardSuit = iota
	diamond
	heart
	spade
)


func compareCards(cardOneVal int, cardOneSuit CardSuit, cardTwoVal int, cardTwoSuit CardSuit) int {
	if cardOneVal < 2 || cardOneVal > 13 || cardTwoVal < 2 || cardTwoVal > 13 {
		fmt.Println("Value is not in range [2,13]")
		return 3
	}
	if (cardOneSuit > 4 || cardOneSuit <0) || (cardTwoSuit > 4 || cardTwoSuit < 0) {
		fmt.Println("Suit is invalid.")
		return 3
	}
	if cardOneVal > cardTwoVal {
		return -1
	}
	if cardOneVal == cardTwoVal {
		if cardOneSuit > cardTwoSuit {
			return -1
		}
		if cardOneSuit < cardTwoSuit {
			return 1
		}
		return 0
	}
	if cardTwoVal > cardOneVal {
		return 1
	}
	return 3
}


func main() {
	result := compareCards(8,2,8,-1)
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