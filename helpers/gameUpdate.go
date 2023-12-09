package helpers

import "gg-buffalo-server/models"

func GameUpdate(GAME *models.GameData) {
	// check for initial start
	if !GAME.Flags.GameActive || GAME.Flags.Resolution {
		if allPlayersReady(GAME) && len(GAME.Players) > 0 {
			GAME.Flags = models.Flags{
				GameActive: true,
				PreGame:    true,
			}
			Deal(GAME)
			resetPlayerReady(GAME)
		}
		return
	}

	// after peeking concludes
	if GAME.Flags.PreGame {
		if allPlayersReady(GAME) {
			hideAllCards(GAME)
			resetPlayerReady(GAME)
			// flag flip
			GAME.Flags.PreGame = false
			GAME.Flags.Drawing = true

			// assign starter and turn pointer
			assignNextStarter(GAME)
			GAME.TurnPointer = GAME.StarterPointer
		}
	}

	if GAME.Flags.Discarded && allPlayersReady(GAME) {
		GAME.Flags = models.Flags{
			GameActive: true,
			Drawing:    true,
		}
		assignTurnPointer(GAME)
	}

	CheckGameEnd(GAME)
}
