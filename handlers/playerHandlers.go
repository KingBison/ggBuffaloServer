package handlers

import (
	"gg-buffalo-server/models"
	"net/http"
)

func HandlePlayerEntry(GAMES *[]models.GameData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()

		gameId := params.Get("gameId")

		name := params.Get("name")
		color := params.Get("color")

		password := params.Get("password")

		if name == "" || color == "" || gameId == "" {
			w.WriteHeader(400)
			w.Write([]byte("error retrieving params"))
			return
		}

		for i, game := range *GAMES {
			if game.GameId == gameId {
				for k, player := range game.Players {
					if player.Name == name {
						(*GAMES)[i].Players[k].Color = color
						w.WriteHeader(200)
						w.Write([]byte("re-entered game"))
						return
					}
				}
				if game.Restricted && game.Password != password {
					w.WriteHeader(404)
					w.Write([]byte("passoword incorrect"))
					return
				}
				(*GAMES)[i].Players = append((*GAMES)[0].Players, models.Player{
					Name:  name,
					Color: color,
					Hand:  []models.Card{},
				})
				w.WriteHeader(200)
				w.Write([]byte("entered game"))
				return
			}
		}

		w.WriteHeader(500)
		w.Write([]byte("game not found"))
	}
}
