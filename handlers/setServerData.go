package handlers

import (
	"encoding/json"
	"gg-buffalo-server/models"
	"io"
	"log"
	"net/http"
)

func SetServerData(GAME *models.GameData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		gameCanidate := models.GameData{}

		if r.URL.Query().Get("name") != "ADMIN" {
			msg := "unauthorized attempt to set game data"
			log.Printf(msg)
			w.WriteHeader(400)
			w.Write([]byte(msg))
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			msg := "error reading body to set server data: " + err.Error()
			log.Printf(msg)
			w.WriteHeader(400)
			w.Write([]byte(msg))
			return
		}

		err = json.Unmarshal(body, &gameCanidate)
		if err != nil {
			msg := "error unmarshaling body to set server data: " + err.Error()
			log.Printf(msg)
			w.WriteHeader(400)
			w.Write([]byte(msg))
			return
		}

		*GAME = gameCanidate

		w.WriteHeader(200)
		w.Write([]byte("update successful"))
	}
}
