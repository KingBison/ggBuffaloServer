package helpers

import (
	"gg-buffalo-server/models"
	"math/rand"
)

var SUITS = []models.Suit{
	{
		Name:  "clubs",
		Value: 0,
	},
	{
		Name:  "diamonds",
		Value: 0,
	},
	{
		Name:  "hearts",
		Value: 0,
	},
	{
		Name:  "spades",
		Value: 0,
	},
}

var NUMBERS = []models.Number{
	{
		Name:  "ace",
		Value: 1,
	},
	{
		Name:  "two",
		Value: 2,
	},
	{
		Name:  "three",
		Value: 3,
	},
	{
		Name:  "four",
		Value: 4,
	},
	{
		Name:  "five",
		Value: 5,
	},
	{
		Name:  "six",
		Value: 6,
	},
	{
		Name:  "seven",
		Value: 7,
	},
	{
		Name:  "eight",
		Value: 8,
	},
	{
		Name:  "nine",
		Value: 9,
	},
	{
		Name:  "ten",
		Value: 10,
	},
	{
		Name:  "jack",
		Value: 15,
	},
	{
		Name:  "queen",
		Value: 15,
	},
	{
		Name:  "king",
		Value: 15,
	},
}

var JOKER = models.Card{
	Suit: models.Suit{
		Name:  "joker",
		Value: 0,
	},
	Number: models.Number{
		Name:  "joker",
		Value: 0,
	},
}

var DECK []models.Card

func init() {
	for _, suit := range SUITS {
		for _, number := range NUMBERS {
			DECK = append(DECK, models.Card{
				Suit:   suit,
				Number: number,
			})
		}
	}

	for i := 0; i < 2; i++ {
		DECK = append(DECK, JOKER)
	}
}

func GetRandomCard(game *models.GameData) models.Card {
	for {
		suggestion := DECK[rand.Intn(len(DECK))]
		if !cardIsInPlay(suggestion, *game) {
			return suggestion
		}
	}
}

func cardIsInPlay(suggestion models.Card, game models.GameData) bool {
	// check for card in any player's hand
	for _, player := range game.Players {
		for _, card := range player.Hand {
			if card.Number.Name == suggestion.Number.Name && card.Suit.Name == suggestion.Suit.Name {
				return true
			}
		}
	}
	// card is on top of discard
	if game.Table.TopOfDiscard.Number.Name == suggestion.Number.Name && game.Table.TopOfDiscard.Suit.Name == suggestion.Suit.Name {
		return true
	}
	return false
}
