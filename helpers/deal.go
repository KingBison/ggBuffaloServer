package helpers

import "gg-buffalo-server/models"

func Deal(GAME *models.GameData) {
	for i := range GAME.Players {
		for k := 0; k < 4; k++ {
			GAME.Players[i].Hand = append(GAME.Players[i].Hand, GetRandomCard(GAME))
		}
	}
}
