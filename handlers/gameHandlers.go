package handlers

import (
	"encoding/json"
	"gg-buffalo-server/models"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type OutgoingGameData struct {
	GameId          string `json:"gameId"`
	Creator         string `json:"creator"`
	NumberOfPlayers int    `json:"numberOfPlayers"`
	Restricted      bool   `json:"restricted"`
}

func GetGames(GAMES *[]models.GameData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		outGoingGameData := []OutgoingGameData{}

		for _, game := range *GAMES {
			outGoingGameData = append(outGoingGameData, OutgoingGameData{
				GameId:          game.GameId,
				Creator:         game.Creator,
				NumberOfPlayers: len(game.Players),
				Restricted:      game.Restricted,
			})
		}

		body, err := json.Marshal(outGoingGameData)
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

func CreateGame(GAMES *[]models.GameData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()

		gameId := uuid.New().String()

		creator := params.Get("creator")
		creatorColor := params.Get("creatorColor")
		password := params.Get("password")

		if creator == "" || creatorColor == "" {
			w.WriteHeader(400)
			w.Write([]byte("error retrieving creator"))
			return
		}

		restrictedBool := false
		if password != "" {
			restrictedBool = true
		}

		newGame := models.GameData{
			GameId:     gameId,
			Creator:    creator,
			Restricted: restrictedBool,
			Password:   password,
			Players: []models.Player{
				{
					Name:  creator,
					Color: creatorColor,
					Hand:  []models.Card{},
				},
			},
		}

		*GAMES = append(*GAMES, newGame)

		w.WriteHeader(201)
		w.Write([]byte(gameId))
	}
}

func DestroyGame(GAMES *[]models.GameData) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()

		gameId := params.Get("gameId")

		name := params.Get("name")
		color := params.Get("color")

		if name == "" || color == "" || gameId == "" {
			w.WriteHeader(400)
			w.Write([]byte("error retrieving params"))
			return
		}

		for i, game := range *GAMES {
			if game.GameId == gameId {
				for _, player := range game.Players {
					if player.Name == game.Creator {
						(*GAMES) = append((*GAMES)[:i], (*GAMES)[i+1:]...)
						w.WriteHeader(200)
						w.Write([]byte("game destroyed"))
					}
				}
			}
			w.WriteHeader(404)
			w.Write([]byte("you are not allowed to destroy this game, only " + game.Creator + " is"))
		}

		w.WriteHeader(500)
		w.Write([]byte("game not found"))
	}
}
