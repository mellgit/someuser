package model

type Message struct {
	Message string `json:"message"`
}

type Error struct {
	Error string `json:"error"`
}

type Active struct {
	Active bool `json:"active"`
}
