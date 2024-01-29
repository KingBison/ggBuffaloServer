package models

type Card struct {
	Suit    Suit   `json:"suit"`
	Number  Number `json:"number"`
	Visible bool   `json:"visible"`
	// Placement Flags
	Swapped bool `json:"swapped"`
	Slammed bool `json:"slammed"`
	// Face Card Flags
	Selected bool `json:"selected"`
}

type Suit struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Number struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
