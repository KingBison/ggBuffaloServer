package helpers

import (
	"gg-buffalo-server/models"
	"time"
)

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

func censorOtherPlayersCards(otherPlayers []models.Player) []models.Player {
	for i := range otherPlayers {
		for k := range otherPlayers[i].Hand {
			otherPlayers[i].Hand[k].Suit = models.Suit{}
			otherPlayers[i].Hand[k].Number = models.Number{}
			otherPlayers[i].Hand[k].Visible = false
		}
	}

	return otherPlayers
}

func showYourPeekedCards(you models.Player) models.Player {
	for i, card := range you.Hand {
		if card.Peeked {
			you.Hand[i].Visible = true
		} else {
			you.Hand[i].Visible = false
		}
	}
	return you
}

func assignYourPeekableCards(you models.Player) models.Player {
	cardsPeeked := 0
	for _, card := range you.Hand {
		if card.Peeked {
			cardsPeeked++
		}
	}
	if cardsPeeked >= 2 {
		return you
	}
	for i, card := range you.Hand {
		if !card.Peeked {
			you.Hand[i].Peekable = true
		} else {
			you.Hand[i].Peekable = false
		}
	}
	return you
}

func assignOtherPlayersPeekedCards(otherPlayers []models.Player) []models.Player {
	for i := range otherPlayers {
		for k, card := range otherPlayers[i].Hand {
			if card.Peeked {
				otherPlayers[i].Hand[k].Peeked = true
			}
		}
	}
	return otherPlayers
}

func assignSlammableCards(you models.Player) models.Player {
	for i, card := range you.Hand {
		if !card.Slammed && !card.FailedSlammed {
			you.Hand[i].Slammable = true
		} else {
			you.Hand[i].Slammable = false
		}
	}

	return you
}

func assignYourFailedSlams(you models.Player) models.Player {
	for i, card := range you.Hand {
		if card.FailedSlammed {
			you.Hand[i].Visible = true
		}
	}

	return you
}

func assignFailedSlams(otherPlayers []models.Player) []models.Player {
	for i := range otherPlayers {
		for k, card := range otherPlayers[i].Hand {
			if card.FailedSlammed {
				otherPlayers[i].Hand[k].Visible = true
			}
		}
	}

	return otherPlayers
}

func assignSwaps(you models.Player) models.Player {
	for i, card := range you.Hand {
		if !card.Slammed && !card.FailedSlammed {
			you.Hand[i].Swappable = true
		}
	}

	return you
}

func assignKingPeekable(you models.Player) models.Player {
	cardsPeeked := 0
	for _, card := range you.Hand {
		if card.KingPeeked {
			cardsPeeked++
		}
	}

	if cardsPeeked != 0 {
		return you
	}

	for i, card := range you.Hand {
		if !card.Slammed && !card.FailedSlammed {
			you.Hand[i].KingPeekable = true
		}
	}

	return you
}

func assignKingPeeked(you models.Player) models.Player {
	for i, card := range you.Hand {
		if card.KingPeeked {
			you.Hand[i].Visible = true
		}
	}

	return you
}

func assignQueenSelectableAndUnSelectable(otherPlayers []models.Player, GAME models.GameData) []models.Player {
	for i, player := range otherPlayers {
		for k, card := range player.Hand {
			if !card.Slammed {
				if !card.QueenSelected {
					if !getQueenSelectCapOut(GAME) {
						otherPlayers[i].Hand[k].QueenSelectable = true
					}
				} else {
					otherPlayers[i].Hand[k].QueenUnSelectable = true
				}
			}
		}
	}

	return otherPlayers
}

func assignYouQueenSelectableAndUnSelectable(you models.Player, GAME models.GameData) models.Player {
	for i, card := range you.Hand {
		if !card.Slammed {
			if !card.QueenSelected {
				if !getQueenSelectCapOut(GAME) {
					you.Hand[i].QueenSelectable = true
				}
			} else {
				you.Hand[i].QueenUnSelectable = true
			}
		}
	}

	return you
}

func assignTurnPointer(otherPlayers []models.Player, GAME models.GameData) []models.Player {
	for i, player := range otherPlayers {
		if player.Name == GAME.Players[GAME.TurnIndex].Name {
			otherPlayers[i].TurnIndicator = true
		}
	}

	return otherPlayers
}

func getQueenSelectCapOut(GAME models.GameData) bool {
	count := 0
	for _, player := range GAME.Players {
		for _, card := range player.Hand {
			if card.QueenSelected {
				count++
			}
		}
	}
	if count >= 2 {
		return true
	} else {
		return false
	}
}

func showAllYourCards(you models.Player) models.Player {
	for i := range you.Hand {
		you.Hand[i].Visible = true
	}

	return you
}

func showAllCards(otherPlayers []models.Player) []models.Player {
	for i := range otherPlayers {
		for k := range otherPlayers[i].Hand {
			otherPlayers[i].Hand[k].Visible = true
		}
	}

	return otherPlayers
}

func assignYourQueenSwapped(you models.Player) models.Player {
	for i := range you.Hand {
		you.Hand[i].QueenSwapped = you.Hand[i].QueenSelected
	}

	return you
}

func assignQueenSwapped(otherPlayers []models.Player) []models.Player {
	for i := range otherPlayers {
		for k := range otherPlayers[i].Hand {
			otherPlayers[i].Hand[k].QueenSwapped = otherPlayers[i].Hand[k].QueenSelected
		}
	}

	return otherPlayers
}

//animations

func assignPeekingAnimations(you models.Player) models.Player {
	currentTime := time.Now()
	for i, card := range you.Hand {
		timeDiff := currentTime.Sub(card.PeekTicker).Seconds()
		if timeDiff < 1.2 {
			you.Hand[i].PeekAni = true
		}
	}
	return you
}

func assignUnPeekingAnimations(you models.Player) models.Player {
	currentTime := time.Now()
	for i, card := range you.Hand {
		timeDiff := currentTime.Sub(card.UnPeekTicker).Seconds()
		if timeDiff < 1.2 {
			you.Hand[i].UnPeekAni = true
		}
	}
	return you
}
