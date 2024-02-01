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

	Swappable bool `json:"swappable"`
	Slammable bool `json:"slammable"`

	Swapped bool `json:"swapped"`
	Slammed bool `json:"slammed"`
	// Face Card Flags
	KingSelectable  bool `json:"kingSelectable"`
	QueenSelectable bool `json:"queenSelectable"`

	KingSelected  bool `json:"kingSelected"`
	QueenSelected bool `json:"queenSelected"`

	QueenUnSelectable bool `json:"queenUnSelected"`
}

type Suit struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Number struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
