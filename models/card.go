package models

type Card struct {
	Suit    Suit   `json:"suit"`
	Number  Number `json:"number"`
	Visible bool   `json:"visible"`
}

type Suit struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Number struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
