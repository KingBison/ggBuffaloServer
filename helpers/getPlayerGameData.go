package helpers

import (
	"errors"
	"gg-buffalo-server/models"
)

func GetPlayerGameData(GAME models.GameData, PLAYER models.Player) (models.OutgoingGameData, error) {
	outgoingOtherPlayers := reorderoutgoingOtherPlayers(GAME.Players, PLAYER)

	if !GAME.Active && !GAME.Resolution {
		// TABLE
		outgoingTable := models.Table{
			TopOfDeck:    models.Card{Visible: false},
			TopOfDiscard: models.Card{Empty: true},
		}

		return models.OutgoingGameData{
			You:          PLAYER,
			OtherPlayers: outgoingOtherPlayers,
			Table:        outgoingTable,
			OtherData:    models.OtherData{},
		}, nil
	}

	if GAME.Peeking {
		outgoingOtherPlayers = censorOtherPlayersCards(outgoingOtherPlayers)
		outgoingOtherPlayers = assignOtherPlayersPeekedCards(outgoingOtherPlayers)

		outgoingTable := models.Table{
			TopOfDeck:    models.Card{Visible: false},
			TopOfDiscard: models.Card{Empty: true},
		}

		outGoingYou := showYourPeekedCards(PLAYER)
		outGoingYou = assignYourPeekableCards(PLAYER)

		return models.OutgoingGameData{
			You:          outGoingYou,
			OtherPlayers: outgoingOtherPlayers,
			Table:        outgoingTable,
			OtherData: models.OtherData{
				BuffaloCalled: false,
			},
		}, nil
	}

	if GAME.Drawing {
		outgoingOtherPlayers = assignFailedSlams(outgoingOtherPlayers)
		outgoingOtherPlayers = assignTurnPointer(outgoingOtherPlayers, GAME)

		outgoingTable := models.Table{
			TopOfDeck: models.Card{Visible: false, Drawable: (GAME.Players[GAME.TurnIndex].Name == PLAYER.Name)},
			TopOfDiscard: models.Card{
				Empty:   GAME.Table.TopOfDiscard.Number.Name == "",
				Visible: true,
				Suit:    GAME.Table.TopOfDiscard.Suit,
				Number:  GAME.Table.TopOfDiscard.Number,
			},
		}

		outGoingYou := assignSlammableCards(PLAYER)
		outGoingYou = assignYourFailedSlams(PLAYER)
		outGoingYou.TurnIndicator = (GAME.Players[GAME.TurnIndex].Name == PLAYER.Name)

		return models.OutgoingGameData{
			You:          outGoingYou,
			OtherPlayers: outgoingOtherPlayers,
			Table:        outgoingTable,
			OtherData: models.OtherData{
				BuffaloCalled:   GAME.OtherData.BuffaloCalled,
				BuffaloCallable: ((GAME.Players[GAME.TurnIndex].Name == PLAYER.Name) && !GAME.OtherData.BuffaloCalled),
				TurnsLeft:       GAME.OtherData.TurnsLeft,
			},
		}, nil

	}

	if GAME.Deciding {
		outgoingOtherPlayers = assignFailedSlams(outgoingOtherPlayers)
		outgoingOtherPlayers = assignTurnPointer(outgoingOtherPlayers, GAME)

		outgoingTable := models.Table{
			TopOfDeck: models.Card{
				Visible: (GAME.Players[GAME.TurnIndex].Name == PLAYER.Name),
				Suit:    GAME.Table.TopOfDeck.Suit,
				Number:  GAME.Table.TopOfDeck.Number,
			},
			TopOfDiscard: models.Card{
				Empty:       GAME.Table.TopOfDiscard.Number.Name == "",
				Visible:     true,
				Suit:        GAME.Table.TopOfDiscard.Suit,
				Number:      GAME.Table.TopOfDiscard.Number,
				Discardable: (GAME.Players[GAME.TurnIndex].Name == PLAYER.Name),
			},
		}

		outGoingYou := assignSlammableCards(PLAYER)
		outGoingYou = assignYourFailedSlams(PLAYER)
		if GAME.Players[GAME.TurnIndex].Name == PLAYER.Name {
			outGoingYou = assignSwaps(PLAYER)
		}
		outGoingYou.TurnIndicator = (GAME.Players[GAME.TurnIndex].Name == PLAYER.Name)

		return models.OutgoingGameData{
			You:          outGoingYou,
			OtherPlayers: outgoingOtherPlayers,
			Table:        outgoingTable,
			OtherData: models.OtherData{
				BuffaloCalled:   GAME.OtherData.BuffaloCalled,
				BuffaloCallable: ((GAME.Players[GAME.TurnIndex].Name == PLAYER.Name) && !GAME.OtherData.BuffaloCalled),
				TurnsLeft:       GAME.OtherData.TurnsLeft,
			},
		}, nil

	}

	if GAME.Discarded {
		outgoingOtherPlayers = assignFailedSlams(outgoingOtherPlayers)
		outgoingOtherPlayers = assignTurnPointer(outgoingOtherPlayers, GAME)

		outgoingTable := models.Table{
			TopOfDeck: models.Card{},
			TopOfDiscard: models.Card{
				Empty:   GAME.Table.TopOfDiscard.Number.Name == "",
				Visible: true,
				Suit:    GAME.Table.TopOfDiscard.Suit,
				Number:  GAME.Table.TopOfDiscard.Number,
			},
		}

		outGoingYou := assignSlammableCards(PLAYER)
		if (GAME.Players[GAME.TurnIndex].Name == PLAYER.Name) && GAME.KingIndicator {
			outGoingYou = assignKingPeekable(PLAYER)
			outGoingYou = assignKingPeeked(PLAYER)
		}
		outGoingYou = assignYourFailedSlams(PLAYER)
		outGoingYou.TurnIndicator = (GAME.Players[GAME.TurnIndex].Name == PLAYER.Name)

		return models.OutgoingGameData{
			You:          outGoingYou,
			OtherPlayers: outgoingOtherPlayers,
			Table:        outgoingTable,
			OtherData: models.OtherData{
				BuffaloCalled:   GAME.OtherData.BuffaloCalled,
				BuffaloCallable: ((GAME.Players[GAME.TurnIndex].Name == PLAYER.Name) && !GAME.OtherData.BuffaloCalled),
				TurnsLeft:       GAME.OtherData.TurnsLeft,
			},
		}, nil

	}

	if GAME.QueenAction {
		outgoingOtherPlayers = assignFailedSlams(outgoingOtherPlayers)
		outgoingOtherPlayers = assignTurnPointer(outgoingOtherPlayers, GAME)

		if GAME.Players[GAME.TurnIndex].Name == PLAYER.Name {
			outgoingOtherPlayers = assignQueenSelectableAndUnSelectable(outgoingOtherPlayers, GAME)
		}

		outgoingTable := models.Table{
			TopOfDeck: models.Card{},
			TopOfDiscard: models.Card{
				Empty:   GAME.Table.TopOfDiscard.Number.Name == "",
				Visible: true,
				Suit:    GAME.Table.TopOfDiscard.Suit,
				Number:  GAME.Table.TopOfDiscard.Number,
			},
		}

		outGoingYou := assignSlammableCards(PLAYER)
		outGoingYou = assignYourFailedSlams(PLAYER)
		outGoingYou.TurnIndicator = (GAME.Players[GAME.TurnIndex].Name == PLAYER.Name)
		if outGoingYou.TurnIndicator {
			outGoingYou = assignYouQueenSelectableAndUnSelectable(outGoingYou, GAME)
		}

		return models.OutgoingGameData{
			You:          outGoingYou,
			OtherPlayers: outgoingOtherPlayers,
			Table:        outgoingTable,
			OtherData: models.OtherData{
				BuffaloCalled:   GAME.OtherData.BuffaloCalled,
				BuffaloCallable: ((GAME.Players[GAME.TurnIndex].Name == PLAYER.Name) && !GAME.OtherData.BuffaloCalled),
				TurnsLeft:       GAME.OtherData.TurnsLeft,
				CanQueenSwap:    (outGoingYou.TurnIndicator && getQueenSelectCapOut(GAME)),
			},
		}, nil

	}

	if GAME.Resolution {
		outgoingOtherPlayers = showAllCards(outgoingOtherPlayers)

		outgoingTable := models.Table{
			TopOfDeck: models.Card{},
			TopOfDiscard: models.Card{
				Empty:   GAME.Table.TopOfDiscard.Number.Name == "",
				Visible: true,
				Suit:    GAME.Table.TopOfDiscard.Suit,
				Number:  GAME.Table.TopOfDiscard.Number,
			},
		}

		outGoingYou := showAllYourCards(PLAYER)

		return models.OutgoingGameData{
			You:          outGoingYou,
			OtherPlayers: outgoingOtherPlayers,
			Table:        outgoingTable,
			OtherData:    models.OtherData{},
		}, nil

	}

	return models.OutgoingGameData{}, errors.New("unknown error")
}
