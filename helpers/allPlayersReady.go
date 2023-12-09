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

func resetPlayerReady(GAME *models.GameData) {
	for i := range GAME.Players {
		GAME.Players[i].Ready = false
	}
}
