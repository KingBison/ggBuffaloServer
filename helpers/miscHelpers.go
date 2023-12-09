package helpers

import "gg-buffalo-server/models"

func totalPeeked(hand []models.Card) int {
	count := 0
	for _, card := range hand {
		if card.Visible {
			count++
		}
	}
	return count
}

func hideAllCards(GAME *models.GameData) {
	for i := range GAME.Players {
		for k := range GAME.Players[i].Hand {
			GAME.Players[i].Hand[k].Visible = false
		}
	}
}

func assignNextStarter(GAME *models.GameData) {
	if GAME.StarterPointer == "" {
		GAME.StarterPointer = GAME.Players[0].Name
		return
	}

	currentStarterI, _ := GetPlayerIndex(GAME.StarterPointer, GAME)

	if currentStarterI+1 == len(GAME.Players) {
		GAME.StarterPointer = GAME.Players[0].Name
	} else {
		GAME.StarterPointer = GAME.Players[currentStarterI+1].Name
	}
}

func assignTurnPointer(GAME *models.GameData) {
	currentTurnI, _ := GetPlayerIndex(GAME.TurnPointer, GAME)
	if currentTurnI+1 == len(GAME.Players) {
		GAME.TurnPointer = GAME.Players[0].Name
	} else {
		GAME.TurnPointer = GAME.Players[currentTurnI+1].Name
	}
}
