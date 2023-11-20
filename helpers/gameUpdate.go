package helpers

import "gg-buffalo-server/models"

func GameUpdate(GAME *models.GameData) {
	// check for initial start
	if !GAME.Flags.GameActive {
		if allPlayersReady(GAME) {
			GAME.Flags.GameActive = true
			GAME.Flags.PreGame = true
			Deal(GAME)
		}
		return
	}
}
