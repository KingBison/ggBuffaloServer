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
			GAME.Players[i].Hand[k] = newCard
		}
	}
}

func playerHasNoCards(GAME *models.GameData) (bool, *models.Player) {
	for i, player := range GAME.Players {
		cardsFound := 0
		for _, card := range player.Hand {
			if !card.Slammed {
				cardsFound++
			}
		}
		if cardsFound == 0 {
			return true, &GAME.Players[i]
		}
	}
	return false, nil
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
