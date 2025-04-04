package models

import (
	"log"
	"placemaking-backend-go/utils"
	"time"
)

type Token struct {
	ID         string    `json:"id"`
	User_id    string    `json:"user_id"`
	Token      string    `json:"token"`
	TokenType  string    `json:"token_type"`
	Active     bool      `json:"active"`
	Created_at string    `json:"created_at"`
	Updated_at string    `json:"updated_at"`
	Expires_at time.Time `json:"expires_at"`
}

// SanitizedToken representa o token sem informações sensíveis.
type SanitizedToken struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
	ExpiresAt string `json:"expires_at"`
	CreatedAt string `json:"created_at"`
}

// SanitizeToken formata os timestamps corretamente e retorna um token seguro.
func SanitizeToken(token Token) SanitizedToken {
	return SanitizedToken{
		Token:     token.Token,
		TokenType: token.TokenType,
		ExpiresAt: token.Expires_at.Format(time.RFC3339),
		CreatedAt: token.Created_at,
	}
}

// Método para converter as strings em time.Time
func (u *Token) ConvertTimestamps() {
	var err error
	u.Created_at, err = utils.ParseTimestamp(u.Created_at)
	if err != nil {
		log.Println("Erro ao converter CreatedAt:", err)
	}

	u.Updated_at, err = utils.ParseTimestamp(u.Updated_at)
	if err != nil {
		log.Println("Erro ao converter UpdatedAt:", err)
	}

	u.Updated_at, err = utils.ParseTimestamp(u.Expires_at.String())
	if err != nil {
		log.Println("Erro ao converter UpdatedAt:", err)
	}
}
