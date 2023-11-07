package main

import (
	"gg-buffalo-server/handlers"
	"gg-buffalo-server/models"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var GAME = &models.GameData{}

func init() {
	GAME.CreatedDate = time.Now().Format(time.RFC3339)
}

func main() {
	log.Println("gray-gaming buffalo server starting up...")

	router := mux.NewRouter()

	router.HandleFunc("/getGameData", handlers.GetServerData(GAME)).Methods("GET")
	router.HandleFunc("/setGameData", handlers.SetServerData(GAME)).Methods("PUT")

	router.HandleFunc("/handlePlayerEntry", handlers.HandlePlayerEntry(GAME)).Methods("GET")
	router.HandleFunc("/handlePlayerExit", handlers.HandlePlayerExit(GAME)).Methods("GET")

	router.Use(handlers.Middleware())

	http.ListenAndServe("127.0.0.1:8080", router)
}
