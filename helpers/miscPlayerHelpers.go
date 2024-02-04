package helpers

import "gg-buffalo-server/models"

func allPlayersReady(GAME *models.GameData) bool {
	for _, player := range GAME.Players {
		if !player.Ready {
			return false
		}
	}
	return true
}

func makeAllPlayersNotReady(GAME *models.GameData) {
	for i := range GAME.Players {
		GAME.Players[i].Ready = false
	}
}

func assignStarterandTurnPointer(GAME *models.GameData) {
	if GAME.StarterIndex == -1 {
		// for the first round, the creator is assigned as the starter
		GAME.StarterIndex = 0
	} else {
		// if the last player is the current starter, make the first palyer the new starter, otherwise make the next player starter
		if GAME.StarterIndex+1 == len(GAME.Players) {
			GAME.StarterIndex = 0
		} else {
			GAME.StarterIndex++
		}
	}

	GAME.TurnIndex = GAME.StarterIndex
}

func assignNextTurn(GAME *models.GameData) {
	if GAME.TurnIndex+1 == len(GAME.Players) {
		GAME.TurnIndex = 0
	} else {
		GAME.TurnIndex++
	}

}
