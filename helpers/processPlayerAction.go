package helpers

import (
	"errors"
	"gg-buffalo-server/models"
	"net/url"
	"strconv"
)

func ProcessPlayerAction(GAME *models.GameData, name string, action string, params url.Values) error {

	playerData, err := GetPlayer(name, GAME)
	if err != nil {
		return err
	}

	if action == "toggle-ready" {
		playerData.Ready = !playerData.Ready
		return nil
	}

	if action == "peek" {
		cardI, err := strconv.Atoi(params.Get("cardI"))
		if err != nil {
			return errors.New("error finding card param to peek: " + err.Error())
		}
		if cardI >= len(playerData.Hand) {
			return errors.New("error finding card: index out of range")
		}
		// begining of game peeking
		if GAME.Flags.PreGame {
			if playerData.Hand[cardI].Visible {
				return errors.New("card has already been peeked")
			}
			if totalPeeked(playerData.Hand) >= 2 {
				return errors.New("you already peeked 2")
			}
			playerData.Hand[cardI].Visible = true
			return nil
		}
		// KING peeking
		if GAME.Flags.KingAction && GAME.Flags.Discarded {

		}
		return errors.New("peeking not allowed right now")
	}

	if action == "draw" {
		if playerData.Name != GAME.TurnPointer {
			return errors.New("NOT YOUR TURN!!!")
		}
		GAME.DrawnCard = GetRandomCard(GAME)
		GAME.Flags.Drawing = false
		GAME.Flags.Deciding = true
		return nil
	}

	return errors.New("request not identified")
}
