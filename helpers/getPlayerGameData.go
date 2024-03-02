package helpers

import (
	"errors"
	"gg-buffalo-server/models"
)

var emptyTable = models.Table{
	TopOfDeck:    models.Card{Visible: false},
	TopOfDiscard: models.Card{Empty: true},
}

func GetPlayerGameData(GAME models.GameData, PLAYER models.Player) (models.OutgoingGameData, error) {
	// ALWAYS reorder other players
	outgoingOtherPlayers := reorderoutgoingOtherPlayers(GAME.Players, PLAYER)
	// bool for turn determination

	// FOR GAME SART
	if !GAME.Active && !GAME.Resolution {
		return models.OutgoingGameData{
			You:          PLAYER,
			OtherPlayers: outgoingOtherPlayers,
			Table:        emptyTable,
			OtherData:    models.OtherData{CanReadyUp: true},
		}, nil
	}

	// FOR PEEKING PHASE
	if GAME.Peeking {
		outgoingOtherPlayers = censorOtherPlayersCards(outgoingOtherPlayers)
		outgoingOtherPlayers = assignOtherPlayersPeekedCards(outgoingOtherPlayers)

		outGoingYou := showYourPeekedCards(PLAYER)
		outGoingYou = assignYourPeekableCards(outGoingYou)
		outGoingYou = assignPeekingAnimations(outGoingYou)

		return models.OutgoingGameData{
			You:          outGoingYou,
			OtherPlayers: outgoingOtherPlayers,
			Table:        emptyTable,
			OtherData:    models.OtherData{CanReadyUp: true},
		}, nil
	}

	// for when you're drawing

	myTurn := (GAME.Players[GAME.TurnIndex].Name == PLAYER.Name)

	if GAME.Drawing {
		outgoingOtherPlayers = assignFailedSlams(outgoingOtherPlayers)
		outgoingOtherPlayers = assignTurnPointer(outgoingOtherPlayers, GAME)

		outGoingYou := assignSlammableCards(PLAYER)
		outGoingYou = assignYourFailedSlams(outGoingYou)
		outGoingYou = assignUnPeekingAnimations(outGoingYou)
		outGoingYou.TurnIndicator = myTurn

		return models.OutgoingGameData{
			You:          outGoingYou,
			OtherPlayers: outgoingOtherPlayers,
			Table: models.Table{
				TopOfDeck: models.Card{Visible: false, Drawable: myTurn},
				TopOfDiscard: models.Card{
					Empty:   GAME.Table.TopOfDiscard.Number.Name == "",
					Visible: true,
					Suit:    GAME.Table.TopOfDiscard.Suit,
					Number:  GAME.Table.TopOfDiscard.Number,
				},
			},
			OtherData: models.OtherData{
				BuffaloCalled:   GAME.OtherData.BuffaloCalled,
				BuffaloCallable: (myTurn && !GAME.OtherData.BuffaloCalled),
				TurnsLeft:       GAME.OtherData.TurnsLeft,
				CanReadyUp:      false,
			},
		}, nil

	}

	if GAME.Deciding {
		outgoingOtherPlayers = assignFailedSlams(outgoingOtherPlayers)
		outgoingOtherPlayers = assignTurnPointer(outgoingOtherPlayers, GAME)

		outGoingYou := assignSlammableCards(PLAYER)
		outGoingYou = assignYourFailedSlams(outGoingYou)
		if myTurn {
			outGoingYou = assignSwaps(PLAYER)
		}
		outGoingYou.TurnIndicator = myTurn

		return models.OutgoingGameData{
			You:          outGoingYou,
			OtherPlayers: outgoingOtherPlayers,
			Table: models.Table{
				TopOfDeck: models.Card{
					Visible: myTurn,
					Suit:    GAME.Table.TopOfDeck.Suit,
					Number:  GAME.Table.TopOfDeck.Number,
				},
				TopOfDiscard: models.Card{
					Empty:       GAME.Table.TopOfDiscard.Number.Name == "",
					Visible:     true,
					Suit:        GAME.Table.TopOfDiscard.Suit,
					Number:      GAME.Table.TopOfDiscard.Number,
					Discardable: myTurn,
				},
			},
			OtherData: models.OtherData{
				BuffaloCalled:   GAME.OtherData.BuffaloCalled,
				BuffaloCallable: (myTurn && !GAME.OtherData.BuffaloCalled),
				TurnsLeft:       GAME.OtherData.TurnsLeft,
				CanReadyUp:      false,
			},
		}, nil

	}

	if GAME.Discarded {
		outgoingOtherPlayers = assignFailedSlams(outgoingOtherPlayers)
		outgoingOtherPlayers = assignTurnPointer(outgoingOtherPlayers, GAME)
		outgoingOtherPlayers = assignQueenSwapped(outgoingOtherPlayers)

		outGoingYou := assignSlammableCards(PLAYER)
		if myTurn && GAME.KingIndicator {
			outGoingYou = assignKingPeekable(outGoingYou)
			outGoingYou = assignKingPeeked(outGoingYou)
		}
		outGoingYou = assignYourQueenSwapped(outGoingYou)
		outGoingYou = assignYourFailedSlams(outGoingYou)
		outGoingYou.TurnIndicator = myTurn

		return models.OutgoingGameData{
			You:          outGoingYou,
			OtherPlayers: outgoingOtherPlayers,
			Table: models.Table{
				TopOfDeck: models.Card{},
				TopOfDiscard: models.Card{
					Empty:   GAME.Table.TopOfDiscard.Number.Name == "",
					Visible: true,
					Suit:    GAME.Table.TopOfDiscard.Suit,
					Number:  GAME.Table.TopOfDiscard.Number,
				},
			},
			OtherData: models.OtherData{
				BuffaloCalled:   GAME.OtherData.BuffaloCalled,
				BuffaloCallable: (myTurn && !GAME.OtherData.BuffaloCalled),
				TurnsLeft:       GAME.OtherData.TurnsLeft,
				CanReadyUp:      true,
			},
		}, nil

	}

	if GAME.QueenAction {
		outgoingOtherPlayers = assignFailedSlams(outgoingOtherPlayers)
		outgoingOtherPlayers = assignTurnPointer(outgoingOtherPlayers, GAME)

		if GAME.Players[GAME.TurnIndex].Name == PLAYER.Name {
			outgoingOtherPlayers = assignQueenSelectableAndUnSelectable(outgoingOtherPlayers, GAME)
		}

		outGoingYou := assignSlammableCards(PLAYER)
		outGoingYou = assignYourFailedSlams(outGoingYou)
		outGoingYou.TurnIndicator = myTurn
		if myTurn {
			outGoingYou = assignYouQueenSelectableAndUnSelectable(outGoingYou, GAME)
		}

		return models.OutgoingGameData{
			You:          outGoingYou,
			OtherPlayers: outgoingOtherPlayers,
			Table: models.Table{
				TopOfDeck: models.Card{},
				TopOfDiscard: models.Card{
					Empty:   GAME.Table.TopOfDiscard.Number.Name == "",
					Visible: true,
					Suit:    GAME.Table.TopOfDiscard.Suit,
					Number:  GAME.Table.TopOfDiscard.Number,
				},
			},
			OtherData: models.OtherData{
				BuffaloCalled:   GAME.OtherData.BuffaloCalled,
				BuffaloCallable: (myTurn && !GAME.OtherData.BuffaloCalled),
				TurnsLeft:       GAME.OtherData.TurnsLeft,
				CanQueenSwap:    (myTurn && getQueenSelectCapOut(GAME)),
				CanReadyUp:      false,
			},
		}, nil

	}

	if GAME.Resolution {
		outgoingOtherPlayers = showAllCards(outgoingOtherPlayers)
		outGoingYou := showAllYourCards(PLAYER)

		return models.OutgoingGameData{
			You:          outGoingYou,
			OtherPlayers: outgoingOtherPlayers,
			Table: models.Table{
				TopOfDeck: models.Card{},
				TopOfDiscard: models.Card{
					Empty:   GAME.Table.TopOfDiscard.Number.Name == "",
					Visible: true,
					Suit:    GAME.Table.TopOfDiscard.Suit,
					Number:  GAME.Table.TopOfDiscard.Number,
				},
			},
			OtherData: models.OtherData{
				CanReadyUp: true,
			},
		}, nil

	}

	return models.OutgoingGameData{}, errors.New("unknown error")
}
