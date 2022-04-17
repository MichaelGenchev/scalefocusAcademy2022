package main

import (

	"github.com/MichaelGenchev/scalefocusAcademy2022/cardgame"
)

func main() {
	
	cardsSlice := make([]cardgame.Card, 0)
	cardsSlice = append(cardsSlice, cardgame.NewCard(8, 1))
	cardsSlice = append(cardsSlice, cardgame.NewCard(9, 1))
	cardsSlice = append(cardsSlice, cardgame.NewCard(10, 2))
	cardsSlice = append(cardsSlice, cardgame.NewCard(12, 1))
	cardsSlice = append(cardsSlice, cardgame.NewCard(12, 2))

	


	cardgame.MaxCard(cardsSlice)
};
