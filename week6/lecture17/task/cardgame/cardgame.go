package cardgame

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







func MaxCard(cards []Card) Card {
	
	max := cards[0]

	for _, c := range cards {
		
		result := CompareCards(max, c)
		if result == 1 {
			max = c
		}
	}
	fmt.Println(max)

	return max 
}




func NewCard(value int, suit CardSuit) Card {
	return Card{value, suit}
}


func CompareCards(cardOne Card, cardTwo Card) int {
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

