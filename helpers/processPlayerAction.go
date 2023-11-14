package helpers

import (
	"errors"
	"gg-buffalo-server/models"
)

func ProcessPlayerAction(GAME *models.GameData, name string, action string) error {

	playerData, err := GetPlayer(name, GAME)
	if err != nil {
		return err
	}

	if action == "toggle-ready" {
		playerData.Ready = !playerData.Ready
		return nil
	}

	return errors.New("request not identified")
}
