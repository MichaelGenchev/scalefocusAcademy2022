package main

import "fmt"

type Card struct {
	value int
	cardSuit int
}

type CardComparator func(cOne Card, cTwo Card) int



func main() {
	cardsSlice := make([]Card, 0)
	cardsSlice = append(cardsSlice, newCard(8, 1))
	cardsSlice = append(cardsSlice, newCard(9, 1))
	cardsSlice = append(cardsSlice, newCard(10, 2))
	cardsSlice = append(cardsSlice, newCard(12, 1))
	cardsSlice = append(cardsSlice, newCard(12, 2))

	maxCard(cardsSlice, compareCards)


	// WITH ANONYMOUS FUNCTION
	maxCard(cardsSlice, func(c1 Card, c2 Card) int{
		if c1.value < 2 || c1.value > 13 || c2.value < 2 || c2.value > 13 {
			fmt.Println("Value is not in range [2,13]")
			return 3
		}
		if (c1.cardSuit > 4 || c1.cardSuit <0) || (c2.cardSuit > 4 || c2.cardSuit < 0) {
			fmt.Println("Suit is invalid.")
			return 3
		}
		if c1.value > c2.value {
			return -1
		}
		if c1.value == c2.value {
			if c1.cardSuit > c2.cardSuit {
				return -1
			}
			if c1.cardSuit < c2.cardSuit {
				return 1
			}
			return 0
		}
		if c2.value > c1.value {
			return 1
		}
		return 3
	})

}
func newCard(value int, suit int) Card {
	return Card{value, suit}
}



func maxCard(cards []Card, comparatorFunc CardComparator) Card {
	max := cards[0]

	for _, c := range cards {
		result := comparatorFunc(max, c) 
		if result == 1 {
			max = c 
		}
	}
	fmt.Println(max)
	return max
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