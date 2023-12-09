package helpers

import "gg-buffalo-server/models"

func CheckGameEnd(GAME *models.GameData) {
	// check for player with 0 cards
	for i, player := range GAME.Players {
		if len(player.Hand) == 0 {
			resetPlayerReady(GAME)
			GAME.Flags = models.Flags{
				GameActive: true,
				Resolution: true,
			}
			GAME.Players[i].Wins++
			return
		}

	}
}
