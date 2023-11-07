package helpers

import (
	"errors"
	"gg-buffalo-server/models"
)

func GetPlayer(name string, GAME *models.GameData) (*models.Player, error) {

	for i, k := range GAME.Players {
		if k.Name == name {
			return &GAME.Players[i], nil
		}
	}

	return &models.Player{}, errors.New("player not found")
}
