package Server

import "time"

type TokenStruct struct {
	Token string `json:"token"`
}

type TokenName struct {
	Name string `json:"name"`
}

type TokenInfoStruct struct {
	TokenStruct
	TokenName
}

type ValidateStruct struct {
	Valid bool `json:"valid"`
}

type Store struct {
	Host      string
	Name      string
	Token     string
	ExpiresAt time.Time
}
