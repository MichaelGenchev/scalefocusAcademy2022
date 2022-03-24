package cardgame

import "testing"


func TestNewCard(t *testing.T) {
	result := NewCard(8, 1)
	testCard := NewCard(8, 1)

	if result != testCard {
		t.Errorf("NewCard function returned %v, want %v", result, testCard)
	}
	
}

func TestMaxCard(t *testing.T) {
	cardsSlice := make([]Card, 0)
	cardsSlice = append(cardsSlice, NewCard(8, 1))
	cardsSlice = append(cardsSlice, NewCard(9, 1))
	cardsSlice = append(cardsSlice, NewCard(10, 2))
	cardsSlice = append(cardsSlice, NewCard(12, 1))
	cardsSlice = append(cardsSlice, NewCard(12, 2))

	result := MaxCard(cardsSlice)

	correctAnswer := NewCard(12, 2)

	if result != correctAnswer {
		t.Errorf("MaxCard function returned %v, want %v", result, correctAnswer)
	}
	
}


func TestCompareCards(t *testing.T) {
	testCases := []struct {
		description string
		cards []Card
		result int
	}{
		{"correctly compares equal cards", []Card{NewCard(8,1), NewCard(8,1)}, 0},
		{"correctly compares different cards (CardTwo is bigger)", []Card{NewCard(8, 1), NewCard(9,1)}, 1},
		{"correctly compares different cards (CardOne is bigger)", []Card{NewCard(8,1), NewCard(7,1)}, -1},
		{"correctly compares different cards (CardOne and CardTwo have equal values but CardOne has bigger suit)", []Card{NewCard(8,3), NewCard(8, 1)}, -1},
		{"Suit is invalid error", []Card{NewCard(8, -1), NewCard(8, 10)}, 3},
		{"Card values are invalid", []Card{NewCard(123, 1), NewCard(29, 2)}, 3},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			cardOne := testCase.cards[0]
			cardTwo := testCase.cards[1]

			if result := CompareCards(cardOne, cardTwo); result != testCase.result {
				t.Errorf("CompareCards function returned %v, want %v", result, testCase.result)
			}
		})
	}
}