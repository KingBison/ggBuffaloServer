package main

import (
	"gg-buffalo-server/handlers"
	"gg-buffalo-server/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var GAMES = &[]models.GameData{}

func main() {
	log.Println("gray-gaming buffalo server starting up...")

	router := mux.NewRouter()
	// admin routes
	router.HandleFunc("/GETGAMES", handlers.GetAllGamesData(GAMES)).Methods("GET")
	router.HandleFunc("/POSTGAMES", handlers.SetGamesData(GAMES)).Methods("GET")

	// game routes
	router.HandleFunc("/getGames", handlers.GetGames(GAMES)).Methods("GET")
	router.HandleFunc("/createGame", handlers.CreateGame(GAMES)).Methods("POST")

	// player routes
	router.HandleFunc("/handlePlayerEntry", handlers.HandlePlayerEntry(GAMES)).Methods("GET")
	router.HandleFunc("/getMyGameData", handlers.GetMyGameData(GAMES)).Methods("GET")

	router.Use(handlers.Middleware())

	http.ListenAndServe("127.0.0.1:8080", router)
}
