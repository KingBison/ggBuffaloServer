package helpers

import (
	"errors"
	"gg-buffalo-server/models"
)

func GetPlayerGameData(GAME models.GameData, PLAYER models.Player) (models.OutgoingGameData, error) {

	if !GAME.Started {
		// OTHERS
		outgoingOtherPlayers := reorderoutgoingOtherPlayers(GAME.Players, PLAYER)
		// TABLE
		outgoingTable := models.Table{
			TopOfDeck:    models.Card{Visible: false},
			TopOfDiscard: models.Card{Empty: true},
		}

		return models.OutgoingGameData{
			You:          PLAYER,
			OtherPlayers: outgoingOtherPlayers,
			Table:        outgoingTable,
			OtherData: models.OtherData{
				BuffaloCalled: false,
			},
		}, nil
	}

	return models.OutgoingGameData{}, errors.New("unknown error")
}

func reorderoutgoingOtherPlayers(players []models.Player, you models.Player) []models.Player {
	index := -1
	for i, player := range players {
		if player.Name == you.Name {
			index = i
			break
		}
	}

	reordered := players[index+1:]
	reordered = append(reordered, players[0:index]...)

	return reordered
}
