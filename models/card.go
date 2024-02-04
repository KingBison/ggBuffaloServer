package models

type Card struct {
	Suit    Suit   `json:"suit"`
	Number  Number `json:"number"`
	Visible bool   `json:"visible"`
	Empty   bool   `json:"empty"`
	// Pre Game Flags
	Peekable bool `json:"peekable"`
	Peeked   bool `json:"peeked"`

	// Placement Flags
	Drawable    bool `json:"drawable"`
	Discardable bool `json:"discardable"`

	Slammable     bool `json:"slammable"`
	Slammed       bool `json:"slammed"`
	FailedSlammed bool `json:"failedSlammed"`

	Swappable bool `json:"swappable"`
	Swapped   bool `json:"swapped"`
	// Face Card Flags
	KingPeekable    bool `json:"kingPeekable"`
	QueenSelectable bool `json:"queenSelectable"`

	KingPeeked    bool `json:"kingPeeked"`
	QueenSelected bool `json:"queenSelected"`

	QueenUnSelectable bool `json:"queenUnSelectable"`
}

type Suit struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Number struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
