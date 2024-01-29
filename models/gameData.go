package models

type GameData struct {
	GameId     string `json:"gameId"`
	Creator    string `json:"creator"`
	Restricted bool   `json:"restricted"`
	Password   string `json:"password"`

	Players        []Player `json:"players"`
	TurnPointer    string   `json:"turnPointer"`
	StarterPointer string   `json:"starterPointer"`

	TopOfDiscard Card `json:"topOfDiscard"`
	DrawnCard    Card `json:"drawnCard"`

	Flags Flags `json:"flags"`
}

type Flags struct {
	// overall game
	GameActive bool `json:"gameActive"`
	PreGame    bool `json:"preGame"`
	Resolution bool `json:"resolution"`

	// while game active
	Drawing       bool `json:"drawing"`
	Deciding      bool `json:"deciding"`
	Discarded     bool `json:"discarded"`
	BuffaloCalled bool `json:"buffaloCalled"`
	TurnsLeft     int  `json:"turnsLeft"`

	// face card flags
	JackAction   bool `json:"jackAction"`
	QueenAction  bool `json:"queenAction"`
	QueenSwapped bool `json:"queenSwapped"`
	KingAction   bool `json:"kingAction"`
}
