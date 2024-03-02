package helpers

import (
	"gg-buffalo-server/models"
)

func UpdateGame(GAME *models.GameData) {

	// GAME START
	if !GAME.Active && allPlayersReady(GAME) {
		handleGameStart(GAME)
		return
	}

	// check for first turn
	if GAME.Peeking && allPlayersReady(GAME) {
		handleFirstTurnAnimations(GAME)
		handleFirstTurn(GAME)
		return
	}

	// check for new turn
	if GAME.Discarded && allPlayersReady(GAME) {
		handleNewTurn(GAME)
	}

	// check for game end
	if GAME.Active && (checkForNoCardsEOG(GAME) || checkForBuffaloCalledEOG(GAME)) {
		handleResolution(GAME)
	}

}
