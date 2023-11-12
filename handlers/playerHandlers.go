package handlers

import (
	"gg-buffalo-server/helpers"
	"gg-buffalo-server/models"
	"log"
	"net/http"
)

func HandlePlayerEntry(GAME *models.GameData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nameParam := r.URL.Query().Get("name")

		player, err := helpers.GetPlayer(nameParam, GAME)
		if err != nil {
			GAME.Players = append(GAME.Players, models.Player{
				Name: nameParam,
				Hand: []models.Card{},
			})
			log.Printf("%s has entered the game", nameParam)
			return
		}

		log.Printf("%s has re-entered the game", player.Name)

	}
}

func HandlePlayerExit(GAME *models.GameData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nameParam := r.URL.Query().Get("name")

		err := helpers.RemovePlayer(nameParam, GAME)
		if err != nil {
			log.Println("ERROR: ", err)
			w.WriteHeader(500)
		}

		log.Printf("%s has exited the game", nameParam)

	}
}
