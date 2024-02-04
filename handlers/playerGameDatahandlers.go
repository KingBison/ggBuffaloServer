package handlers

import (
	"encoding/json"
	"gg-buffalo-server/helpers"
	"gg-buffalo-server/models"
	"net/http"
)

func GetMyGameData(GAMES *[]models.GameData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		params := r.URL.Query()

		gameId := params.Get("gameId")
		name := params.Get("name")

		if name == "" || gameId == "" {
			w.WriteHeader(400)
			w.Write([]byte("error retrieving params"))
			return
		}

		for i, game := range *GAMES {
			if game.GameId == gameId {
				for k, player := range game.Players {
					if player.Name == name {
						GameCopyData, _ := json.Marshal((*GAMES)[i])
						GameCopy := models.GameData{}
						_ = json.Unmarshal(GameCopyData, &GameCopy)

						outgoingGameData, err := helpers.GetPlayerGameData(GameCopy, GameCopy.Players[k])
						if err != nil {
							w.WriteHeader(500)
							w.Write([]byte(err.Error()))
							return
						} else {
							data, err := json.Marshal(outgoingGameData)
							if err != nil {
								w.WriteHeader(500)
								w.Write([]byte("unable to marhal outgoing game data: " + err.Error()))
								return
							}
							w.WriteHeader(200)
							w.Write(data)
							return
						}
					}
				}
				w.WriteHeader(500)
				w.Write([]byte("player not found"))
				return
			}
		}

		w.WriteHeader(500)
		w.Write([]byte("game not found"))
	}
}
