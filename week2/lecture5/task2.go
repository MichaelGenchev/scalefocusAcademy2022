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



func main() {
	cardsSlice := make([]Card, 0)
	cardsSlice = append(cardsSlice, newCard(8, 1))
	cardsSlice = append(cardsSlice, newCard(9, 1))
	cardsSlice = append(cardsSlice, newCard(10, 2))
	cardsSlice = append(cardsSlice, newCard(12, 1))
	cardsSlice = append(cardsSlice, newCard(12, 2))




	maxCard(cardsSlice)


}




func maxCard(cards []Card) Card {
	
	max := cards[0]

	for _, c := range cards {
		
		result := compareCards(max, c)
		if result == 1 {
			max = c
		}
	}
	fmt.Println(max)

	return max 
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

