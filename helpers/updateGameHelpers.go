package helpers

import (
	"gg-buffalo-server/models"
)

func clearAllGameFlags(GAME *models.GameData) {
	GAME.Active = false
	GAME.Peeking = false
	GAME.Drawing = false
	GAME.Deciding = false
	GAME.Discarded = false
	GAME.JackIndicator = false
	GAME.QueenAction = false
	GAME.KingIndicator = false
	GAME.Resolution = false
}

func checkForNoCardsEOG(GAME *models.GameData) bool {
	for i, player := range GAME.Players {
		cardsCounted := 0
		for _, card := range player.Hand {
			if !card.Slammed {
				cardsCounted++
			}
		}
		if cardsCounted == 0 {
			// GAME WIN SCENERIO
			removeAllSlammedCards(GAME)
			GAME.Players[i].Wins++
			return true
		}

	}

	return false
}

func checkForBuffaloCalledEOG(GAME *models.GameData) bool {
	if GAME.OtherData.BuffaloCalled && GAME.OtherData.TurnsLeft == 0 {
		// GAME WIN SCENERIO
		eogPlayers := []eogPlayer{}

		for i, player := range GAME.Players {
			eogPlayers = append(eogPlayers, eogPlayer{
				playerIndex: i,
				mainScore:   calculateMainScore(player.Hand),
				highestCard: calculateHighestCard(player.Hand),
			})
		}

		WINNERS := []*models.Player{}

		bestScore := 10000
		highestCardWithBestScore := -1

		for _, player := range eogPlayers {
			if player.mainScore < bestScore {
				WINNERS = []*models.Player{&(*&GAME.Players[player.playerIndex])}
				bestScore = player.mainScore
			} else if player.mainScore == bestScore {
				if player.highestCard > highestCardWithBestScore {
					WINNERS = []*models.Player{&(*&GAME.Players[player.playerIndex])}
					highestCardWithBestScore = player.highestCard
				} else if player.highestCard == highestCardWithBestScore {
					WINNERS = append(WINNERS, &(*&GAME.Players[player.playerIndex]))
				}
			}
		}
		for _, winner := range WINNERS {
			winner.Wins++
		}
		return true
	}
	return false
}

type eogPlayer struct {
	playerIndex int
	mainScore   int
	highestCard int
}

func calculateMainScore(hand []models.Card) int {
	total := 0
	for _, card := range hand {
		total += card.Number.Value
	}
	return total
}

func calculateHighestCard(hand []models.Card) int {
	highestCard := -1
	for _, card := range hand {
		if card.Number.Value > highestCard {
			highestCard = card.Number.Value
		}
	}
	return highestCard
}
