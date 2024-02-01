package handlers

import (
	"encoding/json"
	"gg-buffalo-server/models"
	"io"
	"log"
	"net/http"
)

func GetAllGamesData(GAMES *[]models.GameData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := json.Marshal(*GAMES)
		if err != nil {
			log.Printf("error marshaling request body: %s", err)
			w.WriteHeader(500)
			w.Write(body)
			return
		}
		w.WriteHeader(200)
		w.Write(body)
	}
}

func SetGamesData(GAMES *[]models.GameData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("error reading request body: %s", err)
			w.WriteHeader(500)
			return
		}

		newData := []models.GameData{}

		err = json.Unmarshal(body, &newData)
		if err != nil {
			log.Printf("error unmarshaling request body: %s", err)
			w.WriteHeader(500)
			return
		}

		*GAMES = newData

		w.WriteHeader(200)
		w.Write(body)
	}
}
