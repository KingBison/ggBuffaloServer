package handlers

import (
	"encoding/json"
	"gg-buffalo-server/models"
	"log"
	"net/http"
)

func GetServerData(GAME *models.GameData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := json.Marshal(*GAME)
		if err != nil {
			log.Printf("error marshaling request body: %s", err)
			w.WriteHeader(500)
			w.Write(body)
		}

		w.WriteHeader(200)
		w.Write(body)
	}
}
