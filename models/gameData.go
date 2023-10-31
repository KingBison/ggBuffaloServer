package models

type GameData struct {
	CreatedDate string `json:"createdDate"`

	Players     []Player `json:"players"`
	TurnPointer string   `json:"turnPointer"`

	Flags Flags `json:"flags"`
}

type Flags struct {
	// overall game
	GameActive bool `json:"gameActive"`
	PreGame    bool `json:"preGame"`

	// while game active
	Drawing       bool `json:"drawing"`
	Deciding      bool `json:"deciding"`
	Discarded     bool `json:"discarded"`
	BuffaloCalled bool `json:"buffaloCalled"`

	// face card flags
	JackAction  bool `json:"jackAction"`
	QueenAction bool `json:"queenAction"`
	KingAction  bool `json:"kingAction"`
}
