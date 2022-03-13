package carddraw

import (
	// "fmt"

	"github.com/MichaelGenchev/scalefocusAcademy2022/cardgame"
)

type dealer interface {
	Deal() (*cardgame.Card, error)
	Done() bool
}

func DrawAllCards(dealer dealer) ([]cardgame.Card, error){

	// call the dealer's Draw() method, until you reach a nil Card
	cards := make([]cardgame.Card, 0)

	// for i := 0; i < 52; i++ {
	// 	cards = append(cards, *dealer.Deal())

	// }
	for {
		card, err := dealer.Deal()
		if err != nil {
			if dealer.Done() == true{
				return cards, nil
			}else{
				return nil, err
			}
		}
		cards = append(cards, *card)
	}
}