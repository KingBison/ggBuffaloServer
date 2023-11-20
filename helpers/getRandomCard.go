package helpers

import (
	"gg-buffalo-server/models"
	"math/rand"
)

var DECK = []models.Card{
	{
		Suit: models.Suit{
			Name:  "clubs",
			Value: 0,
		},
		Number: models.Number{
			Name:  "ace",
			Value: 1,
		},
	},
	{
		Suit: models.Suit{
			Name:  "clubs",
			Value: 0,
		},
		Number: models.Number{
			Name:  "two",
			Value: 2,
		},
	},
	{
		Suit: models.Suit{
			Name:  "clubs",
			Value: 0,
		},
		Number: models.Number{
			Name:  "three",
			Value: 3,
		},
	},
	{
		Suit: models.Suit{
			Name:  "clubs",
			Value: 0,
		},
		Number: models.Number{
			Name:  "four",
			Value: 4,
		},
	},
	{
		Suit: models.Suit{
			Name:  "clubs",
			Value: 0,
		},
		Number: models.Number{
			Name:  "five",
			Value: 5,
		},
	},
	{
		Suit: models.Suit{
			Name:  "clubs",
			Value: 0,
		},
		Number: models.Number{
			Name:  "six",
			Value: 6,
		},
	},
	{
		Suit: models.Suit{
			Name:  "clubs",
			Value: 0,
		},
		Number: models.Number{
			Name:  "seven",
			Value: 7,
		},
	},
	{
		Suit: models.Suit{
			Name:  "clubs",
			Value: 0,
		},
		Number: models.Number{
			Name:  "eight",
			Value: 8,
		},
	},
	{
		Suit: models.Suit{
			Name:  "clubs",
			Value: 0,
		},
		Number: models.Number{
			Name:  "nine",
			Value: 9,
		},
	},
	{
		Suit: models.Suit{
			Name:  "clubs",
			Value: 0,
		},
		Number: models.Number{
			Name:  "ten",
			Value: 10,
		},
	},
	{
		Suit: models.Suit{
			Name:  "clubs",
			Value: 0,
		},
		Number: models.Number{
			Name:  "jack",
			Value: 15,
		},
	},
	{
		Suit: models.Suit{
			Name:  "clubs",
			Value: 0,
		},
		Number: models.Number{
			Name:  "queen",
			Value: 15,
		},
	},
	{
		Suit: models.Suit{
			Name:  "clubs",
			Value: 0,
		},
		Number: models.Number{
			Name:  "king",
			Value: 15,
		},
	},
	{
		Suit: models.Suit{
			Name:  "diamonds",
			Value: 1,
		},
		Number: models.Number{
			Name:  "ace",
			Value: 1,
		},
	},
	{
		Suit: models.Suit{
			Name:  "diamonds",
			Value: 1,
		},
		Number: models.Number{
			Name:  "two",
			Value: 2,
		},
	},
	{
		Suit: models.Suit{
			Name:  "diamonds",
			Value: 1,
		},
		Number: models.Number{
			Name:  "three",
			Value: 3,
		},
	},
	{
		Suit: models.Suit{
			Name:  "diamonds",
			Value: 1,
		},
		Number: models.Number{
			Name:  "four",
			Value: 4,
		},
	},
	{
		Suit: models.Suit{
			Name:  "diamonds",
			Value: 1,
		},
		Number: models.Number{
			Name:  "five",
			Value: 5,
		},
	},
	{
		Suit: models.Suit{
			Name:  "diamonds",
			Value: 1,
		},
		Number: models.Number{
			Name:  "six",
			Value: 6,
		},
	},
	{
		Suit: models.Suit{
			Name:  "diamonds",
			Value: 1,
		},
		Number: models.Number{
			Name:  "seven",
			Value: 7,
		},
	},
	{
		Suit: models.Suit{
			Name:  "diamonds",
			Value: 1,
		},
		Number: models.Number{
			Name:  "eight",
			Value: 8,
		},
	},
	{
		Suit: models.Suit{
			Name:  "diamonds",
			Value: 1,
		},
		Number: models.Number{
			Name:  "nine",
			Value: 9,
		},
	},
	{
		Suit: models.Suit{
			Name:  "diamonds",
			Value: 1,
		},
		Number: models.Number{
			Name:  "ten",
			Value: 10,
		},
	},
	{
		Suit: models.Suit{
			Name:  "diamonds",
			Value: 1,
		},
		Number: models.Number{
			Name:  "jack",
			Value: 15,
		},
	},
	{
		Suit: models.Suit{
			Name:  "diamonds",
			Value: 1,
		},
		Number: models.Number{
			Name:  "queen",
			Value: 15,
		},
	},
	{
		Suit: models.Suit{
			Name:  "diamonds",
			Value: 1,
		},
		Number: models.Number{
			Name:  "king",
			Value: 15,
		},
	},
	{
		Suit: models.Suit{
			Name:  "hearts",
			Value: 2,
		},
		Number: models.Number{
			Name:  "ace",
			Value: 1,
		},
	},
	{
		Suit: models.Suit{
			Name:  "hearts",
			Value: 2,
		},
		Number: models.Number{
			Name:  "two",
			Value: 2,
		},
	},
	{
		Suit: models.Suit{
			Name:  "hearts",
			Value: 2,
		},
		Number: models.Number{
			Name:  "three",
			Value: 3,
		},
	},
	{
		Suit: models.Suit{
			Name:  "hearts",
			Value: 2,
		},
		Number: models.Number{
			Name:  "four",
			Value: 4,
		},
	},
	{
		Suit: models.Suit{
			Name:  "hearts",
			Value: 2,
		},
		Number: models.Number{
			Name:  "five",
			Value: 5,
		},
	},
	{
		Suit: models.Suit{
			Name:  "hearts",
			Value: 2,
		},
		Number: models.Number{
			Name:  "six",
			Value: 6,
		},
	},
	{
		Suit: models.Suit{
			Name:  "hearts",
			Value: 2,
		},
		Number: models.Number{
			Name:  "seven",
			Value: 7,
		},
	},
	{
		Suit: models.Suit{
			Name:  "hearts",
			Value: 2,
		},
		Number: models.Number{
			Name:  "eight",
			Value: 8,
		},
	},
	{
		Suit: models.Suit{
			Name:  "hearts",
			Value: 2,
		},
		Number: models.Number{
			Name:  "nine",
			Value: 9,
		},
	},
	{
		Suit: models.Suit{
			Name:  "hearts",
			Value: 2,
		},
		Number: models.Number{
			Name:  "ten",
			Value: 10,
		},
	},
	{
		Suit: models.Suit{
			Name:  "hearts",
			Value: 2,
		},
		Number: models.Number{
			Name:  "jack",
			Value: 15,
		},
	},
	{
		Suit: models.Suit{
			Name:  "hearts",
			Value: 2,
		},
		Number: models.Number{
			Name:  "queen",
			Value: 15,
		},
	},
	{
		Suit: models.Suit{
			Name:  "hearts",
			Value: 2,
		},
		Number: models.Number{
			Name:  "king",
			Value: 15,
		},
	},
	{
		Suit: models.Suit{
			Name:  "spades",
			Value: 3,
		},
		Number: models.Number{
			Name:  "ace",
			Value: 1,
		},
	},
	{
		Suit: models.Suit{
			Name:  "spades",
			Value: 3,
		},
		Number: models.Number{
			Name:  "two",
			Value: 2,
		},
	},
	{
		Suit: models.Suit{
			Name:  "spades",
			Value: 3,
		},
		Number: models.Number{
			Name:  "three",
			Value: 3,
		},
	},
	{
		Suit: models.Suit{
			Name:  "spades",
			Value: 3,
		},
		Number: models.Number{
			Name:  "four",
			Value: 4,
		},
	},
	{
		Suit: models.Suit{
			Name:  "spades",
			Value: 3,
		},
		Number: models.Number{
			Name:  "five",
			Value: 5,
		},
	},
	{
		Suit: models.Suit{
			Name:  "spades",
			Value: 3,
		},
		Number: models.Number{
			Name:  "six",
			Value: 6,
		},
	},
	{
		Suit: models.Suit{
			Name:  "spades",
			Value: 3,
		},
		Number: models.Number{
			Name:  "seven",
			Value: 7,
		},
	},
	{
		Suit: models.Suit{
			Name:  "spades",
			Value: 3,
		},
		Number: models.Number{
			Name:  "eight",
			Value: 8,
		},
	},
	{
		Suit: models.Suit{
			Name:  "spades",
			Value: 3,
		},
		Number: models.Number{
			Name:  "nine",
			Value: 9,
		},
	},
	{
		Suit: models.Suit{
			Name:  "spades",
			Value: 3,
		},
		Number: models.Number{
			Name:  "ten",
			Value: 10,
		},
	},
	{
		Suit: models.Suit{
			Name:  "spades",
			Value: 3,
		},
		Number: models.Number{
			Name:  "jack",
			Value: 15,
		},
	},
	{
		Suit: models.Suit{
			Name:  "spades",
			Value: 3,
		},
		Number: models.Number{
			Name:  "queen",
			Value: 15,
		},
	},
	{
		Suit: models.Suit{
			Name:  "spades",
			Value: 3,
		},
		Number: models.Number{
			Name:  "king",
			Value: 15,
		},
	},
	{
		Suit: models.Suit{
			Name:  "joker",
			Value: 4,
		},
		Number: models.Number{
			Name:  "joker",
			Value: 0,
		},
	},
	{
		Suit: models.Suit{
			Name:  "joker",
			Value: 4,
		},
		Number: models.Number{
			Name:  "joker",
			Value: 0,
		},
	},
}

func GetRandomCard(game *models.GameData) models.Card {
	deckLength := len(DECK)
	for {
		randI := rand.Intn(deckLength)
		suggestion := DECK[randI]
		found := false
		for _, player := range game.Players {
			for _, card := range player.Hand {
				if card.Number.Name == suggestion.Number.Name && card.Suit.Name == suggestion.Suit.Name {
					found = true
					break
				}
			}
			if found == true {
				break
			}
		}
		if game.TopOfDiscard.Number.Name == suggestion.Number.Name && game.TopOfDiscard.Suit.Name == suggestion.Suit.Name {
			found = true
		}

		if !found {
			return suggestion
		}
	}
}
