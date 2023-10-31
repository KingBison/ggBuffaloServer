package models

type Player struct {
	Name string `json:"name"`

	Wins int `json:"wins"`

	Hand []Card `json:"hand"`

	Ready bool `json:"ready"`
}
