package helpers

import "gg-buffalo-server/models"

func handleGameStart(GAME *models.GameData) {
	// clear all flags from previous and deal
	clearAllGameFlags(GAME)
	clearAllCardFlags(GAME)
	clearHandsAndDeal(GAME)
	// reset other data
	GAME.OtherData = models.OtherData{TurnsLeft: -1}
	// clear table
	GAME.Table = models.Table{
		TopOfDeck:    models.Card{Empty: true},
		TopOfDiscard: models.Card{},
	}
	// flip peeking flags and
	GAME.Peeking = true
	GAME.Active = true
	makeAllPlayersNotReady(GAME)
}

func handleFirstTurn(GAME *models.GameData) {
	// clear peeked cards
	clearAllCardFlags(GAME)

	assignStarterandTurnPointer(GAME)

	GAME.Peeking = false
	GAME.Drawing = true

	makeAllPlayersNotReady(GAME)
}

func handleNewTurn(GAME *models.GameData) {
	GAME.Discarded = false
	removeAllSlammedCards(GAME)
	clearAllCardFlags(GAME)
	GAME.KingIndicator = false
	if !GAME.JackIndicator {
		assignNextTurn(GAME)
		if GAME.OtherData.BuffaloCalled {
			GAME.OtherData.TurnsLeft--
		}
	} else {
		GAME.JackIndicator = false
	}
	GAME.Drawing = true
	makeAllPlayersNotReady(GAME)
}

func handleResolution(GAME *models.GameData) {
	clearAllCardFlags(GAME)
	clearAllGameFlags(GAME)
	GAME.TurnIndex = -1
	GAME.Resolution = true
}
