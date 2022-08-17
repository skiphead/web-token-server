package Server

import "time"

type TokenStruct struct {
	SessionToken string `json:"session_token"`
	RefreshToken string `json:"refresh_token"`
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

type SessionTokenStorageStruct struct {
	Host      string
	Name      string
	Token     string
	ExpiresAt time.Time
}

type RefreshTokenStorageStruct struct {
	NameSessionToken string    `json:"name_session_token,omitempty"`
	RefreshToken     string    `json:"refresh_token,omitempty"`
	ExpiresAt        time.Time `json:"expires_at"`
}
