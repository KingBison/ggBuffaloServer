package helpers

import (
	"errors"
	"gg-buffalo-server/models"
	"net/url"
	"strconv"
)

func HandleRequest(GAME *models.GameData, PLAYER *models.Player, PARAMS url.Values) error {

	requestedAction := PARAMS.Get("action")

	if requestedAction == "" {
		return errors.New("no action in params")
	}

	if requestedAction == "toggle-ready" {
		PLAYER.Ready = !PLAYER.Ready
		return nil
	}

	if requestedAction == "pre-game-peek" {
		cardI, err := strconv.Atoi(PARAMS.Get("cardI"))
		if err != nil {
			return errors.New("error parsing cardI")
		}
		if cardI >= len(PLAYER.Hand) {
			return errors.New("indexing error")
		}
		if PLAYER.Hand[cardI].Peeked {
			return errors.New("this card has already been peeked")
		}
		totalPeeked := 0
		for _, card := range PLAYER.Hand {
			if card.Peeked {
				totalPeeked++
			}
		}
		if totalPeeked >= 2 {
			return errors.New("2 cards have already been peeked")
		}
		PLAYER.Hand[cardI].Peeked = true
		return nil
	}

	if requestedAction == "call-buffalo" {
		GAME.OtherData.BuffaloCalled = true
		GAME.OtherData.TurnsLeft = len(GAME.Players) + 1
		return nil
	}

	if requestedAction == "SLAM" {
		cardI, err := strconv.Atoi(PARAMS.Get("cardI"))
		if err != nil {
			return errors.New("error parsing cardI")
		}
		if cardI >= len(PLAYER.Hand) {
			return errors.New("indexing error")
		}
		if PLAYER.Hand[cardI].Number.Name == GAME.Table.TopOfDiscard.Number.Name {
			PLAYER.Hand[cardI].Slammed = true
		} else {
			PLAYER.Hand[cardI].FailedSlammed = true
			PLAYER.Hand = append(PLAYER.Hand, GetRandomCard(GAME))
		}

		return nil
	}

	if requestedAction == "draw" {
		GAME.Table.TopOfDeck = GetRandomCard(GAME)
		GAME.Drawing = false
		GAME.Deciding = true
		return nil
	}

	if requestedAction == "discard" {
		GAME.Table.TopOfDiscard = GAME.Table.TopOfDeck
		GAME.Table.TopOfDeck = models.Card{}
		GAME.Deciding = false
		if GAME.Table.TopOfDiscard.Number.Name == "queen" {
			GAME.QueenAction = true
		} else {
			GAME.Discarded = true
		}
		return nil
	}

	if requestedAction == "swap" {
		cardI, err := strconv.Atoi(PARAMS.Get("cardI"))
		if err != nil {
			return errors.New("error parsing cardI")
		}
		if cardI >= len(PLAYER.Hand) {
			return errors.New("indexing error")
		}
		GAME.Table.TopOfDiscard = PLAYER.Hand[cardI]
		PLAYER.Hand[cardI] = GAME.Table.TopOfDeck
		GAME.Deciding = false
		assignFaceCardFlags(GAME)
		return nil
	}

	if requestedAction == "king-peek" {
		cardI, err := strconv.Atoi(PARAMS.Get("cardI"))
		if err != nil {
			return errors.New("error parsing cardI")
		}
		if cardI >= len(PLAYER.Hand) {
			return errors.New("indexing error")
		}
		totalPeeked := 0
		for _, card := range PLAYER.Hand {
			if card.KingPeeked {
				totalPeeked++
			}
		}
		if totalPeeked == 0 {
			PLAYER.Hand[cardI].KingPeeked = true
		} else {
			return errors.New("you already peeked one")
		}
		return nil
	}

	if requestedAction == "queen-toggle-select" {
		cardI, err := strconv.Atoi(PARAMS.Get("cardI"))
		if err != nil {
			return errors.New("error parsing cardI")
		}
		OPname := PARAMS.Get("OPname")
		if OPname == "" {
			return errors.New("error getting other player's name")
		}
		for i, player := range GAME.Players {
			if player.Name == OPname {
				if cardI >= len(player.Hand) {
					return errors.New("indexing error")
				}
				GAME.Players[i].Hand[cardI].QueenSelected = !GAME.Players[i].Hand[cardI].QueenSelected
			}
		}
		return nil
	}

	if requestedAction == "queen-submit" {
		cardsSelected := []CardSelected{}
		for i, player := range GAME.Players {
			for k, card := range player.Hand {
				if card.QueenSelected {
					cardsSelected = append(cardsSelected, CardSelected{
						PlayerIndex: i,
						CardIndex:   k,
					})
				}
			}
		}
		if len(cardsSelected) != 2 {
			return errors.New("you need to have exactly 2 cards selected")
		}
		temp := GAME.Players[cardsSelected[0].PlayerIndex].Hand[cardsSelected[0].CardIndex]
		GAME.Players[cardsSelected[0].PlayerIndex].Hand[cardsSelected[0].CardIndex] = GAME.Players[cardsSelected[1].PlayerIndex].Hand[cardsSelected[1].CardIndex]
		GAME.Players[cardsSelected[1].PlayerIndex].Hand[cardsSelected[1].CardIndex] = temp

		GAME.QueenAction = false
		GAME.Discarded = true
		return nil
	}

	return errors.New("request not identified")
}

func assignFaceCardFlags(GAME *models.GameData) {
	if GAME.Table.TopOfDiscard.Number.Name == "jack" {
		GAME.JackIndicator = true
		GAME.Discarded = true
	} else if GAME.Table.TopOfDiscard.Number.Name == "queen" {
		GAME.QueenAction = true
	} else if GAME.Table.TopOfDiscard.Number.Name == "king" {
		GAME.KingIndicator = true
		GAME.Discarded = true
	} else {
		GAME.Discarded = true
	}
}

type CardSelected struct {
	PlayerIndex int
	CardIndex   int
}
