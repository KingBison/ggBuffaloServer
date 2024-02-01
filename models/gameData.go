package models

type GameData struct {
	// Base Components
	GameId     string `json:"gameId"`
	Creator    string `json:"creator"`
	Restricted bool   `json:"restricted"`
	Password   string `json:"password"`

	// Players
	Players []Player `json:"players"`

	// Board Values
	Table Table `json:"table"`

	// Phase Values
	Started bool `json:"started"`
	Active  bool `json:"active"`
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
	BuffaloCalled bool `json:"buffaloCalled"`
	TurnsLeft     int  `json:"turnsLeft"`
}
