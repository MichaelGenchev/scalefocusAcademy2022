package carddraw

import (
	// "fmt"

	"github.com/MichaelGenchev/scalefocusAcademy2022/cardgame"
)

type dealer interface {
	Deal() *cardgame.Card
}

func DrawAllCards(dealer dealer) []cardgame.Card {

	// call the dealer's Draw() method, until you reach a nil Card
	cards := make([]cardgame.Card, 0)

	// for i := 0; i < 52; i++ {
	// 	cards = append(cards, *dealer.Deal())

	// }
	for {
		card := dealer.Deal()
		if card == nil{
			break
		}
		cards = append(cards, *card)
	}



	return cards
}


