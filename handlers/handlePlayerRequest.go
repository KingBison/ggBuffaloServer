package handlers

import (
	"gg-buffalo-server/helpers"
	"gg-buffalo-server/models"
	"net/http"
)

func HandlePlayerRequest(GAME *models.GameData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()

		name := params.Get("name")
		action := params.Get("action")

		if name == "" || action == "" {
			w.WriteHeader(400)
			w.Write([]byte("player name or action missing"))
			return
		}

		err := helpers.ProcessPlayerAction(GAME, name, action, params)

		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("request accepted"))
	}
}
