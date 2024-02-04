package models

type GameData struct {
	// Base Components
	GameId     string `json:"gameId"`
	Creator    string `json:"creator"`
	Restricted bool   `json:"restricted"`
	Password   string `json:"password"`

	// Players
	Players []Player `json:"players"`

	StarterIndex int `json:"starterIndex"`
	TurnIndex    int `json:"turnIndex"`

	// Board Values
	Table Table `json:"table"`

	// Phase Values
	Active        bool `json:"active"`
	Peeking       bool `json:"peeking"`
	Drawing       bool `json:"drawing"`
	Deciding      bool `json:"deciding"`
	Discarded     bool `json:"discarded"`
	QueenAction   bool `json:"queenAction"`
	JackIndicator bool `json:"jackIndicator"`
	KingIndicator bool `json:"kingIndicator"`

	Resolution bool `json:"resolution"`

	OtherData OtherData `json:"otherData"`
}

type OutgoingGameData struct {
	You          Player    `json:"you"`
	OtherPlayers []Player  `json:"otherPlayers"`
	Table        Table     `json:"table"`
	OtherData    OtherData `json:"otherData"`
}

type Table struct {
	TopOfDeck    Card `json:"topOfDeck"`
	TopOfDiscard Card `json:"topOfDiscard"`
}

type OtherData struct {
	BuffaloCalled   bool `json:"buffaloCalled"`
	BuffaloCallable bool `json:"buffaloCallable"`
	TurnsLeft       int  `json:"turnsLeft"`
	CanQueenSwap    bool `json:"canQueenSwap"`
}
