package cardgame





var suit = [4]string{"Hearts", "Diamonds", "Clubs", "Spades"}

var face = [13]string{"Ace", "Deuce", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

const rows = 13
const columns = 4
const total = rows * columns

func newCard(face string, cardSuit string) Card {
	return Card{face, cardSuit}
}

type Card struct {
	face string
	cardSuit string
}
func New() Deck {
	var cards []Card = initializeDeck()
	return Deck{cards}
}

type Deck struct {
	cards []Card
}

func initializeDeck() []Card {
	var initDeck []Card

	for r := 0; r <= rows-1; r++ {
		for c := 0; c <= columns-1; c++ {
				face := face[r] 
				suit := suit[c]
				card := newCard(face, suit)
				initDeck = append(initDeck, card)
		}
	}
	return initDeck
}


func (d *Deck) Deal() *Card {
	if len(d.cards) == 0 {
		return &Card{}
	}
	firstCard := d.cards[0]
	d.cards = append(d.cards[:0], d.cards[1:]...)
	return &firstCard
}

