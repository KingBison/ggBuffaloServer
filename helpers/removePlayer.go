package helpers

import (
	"errors"
	"gg-buffalo-server/models"
)

func RemovePlayer(name string, GAME *models.GameData) error {
	for i, k := range GAME.Players {
		if k.Name == name {
			GAME.Players = append(GAME.Players[:i], GAME.Players[i+1:]...)
			return nil
		}
	}

	return errors.New("player not found")
}
