package helpers

import (
	"gg-buffalo-server/models"
)

func clearHandsAndDeal(GAME *models.GameData) {
	for i := range GAME.Players {
		GAME.Players[i].Hand = []models.Card{}
		for k := 0; k < 4; k++ {
			GAME.Players[i].Hand = append(GAME.Players[i].Hand, GetRandomCard(GAME))
		}
	}
}

func clearAllCardFlags(GAME *models.GameData) {
	for i := range GAME.Players {
		for k, card := range GAME.Players[i].Hand {
			newCard := models.Card{}
			newCard.Suit = card.Suit
			newCard.Number = card.Number
			newCard.UnPeekTicker = card.UnPeekTicker
			GAME.Players[i].Hand[k] = newCard
		}
	}
}

func removeAllSlammedCards(GAME *models.GameData) {
	for i, player := range GAME.Players {
		newHand := []models.Card{}
		for _, card := range player.Hand {
			if !card.Slammed {
				newHand = append(newHand, card)
			}
		}
		GAME.Players[i].Hand = newHand

	}
}
