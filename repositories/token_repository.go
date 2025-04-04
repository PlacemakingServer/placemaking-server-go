package repository

import (
	"log"
	"placemaking-backend-go/db"
	"placemaking-backend-go/models"
	"time"
)

// GetToken busca um token específico no banco de dados
func GetToken(token string) (*models.Token, error) {
	supabase := db.GetSupabase()

	var tokens []models.Token
	_, err := supabase.From("tokens").
		Select("*", "", false).
		Eq("token", token).
		ExecuteTo(&tokens)

	if err != nil {
		log.Println("[GetToken] Erro:", err)
		return nil, err
	}

	if len(tokens) > 0 {
		return &tokens[0], nil
	}
	return nil, nil
}

// GetTokenByUserID busca um token pelo user_id e tipo
func GetTokenByUserID(userID string, tokenType string) (*models.Token, error) {
	supabase := db.GetSupabase()

	var tokens []models.Token
	_, err := supabase.From("tokens").
		Select("*", "", false).
		Eq("user_id", userID).
		Eq("token_type", tokenType).
		ExecuteTo(&tokens)

	if err != nil {
		log.Println("[GetTokenByUserID] Erro:", err)
		return nil, err
	}

	if len(tokens) > 0 {
		return &tokens[0], nil
	}
	return nil, nil
}

// InsertToken insere um novo token no banco
func InsertToken(userID string, token string, tokenType string, expiresAt time.Time) (*models.Token, error) {
	supabase := db.GetSupabase()

	newToken := models.Token{
		User_id:    userID,
		Token:      token,
		TokenType:  tokenType,
		Expires_at: expiresAt.Format(time.RFC3339), // Garante o formato correto
		Created_at: time.Now().Format(time.RFC3339),
		Active:     true,
	}

	var insertedTokens []models.Token
	_, err := supabase.From("tokens").
		Insert(map[string]interface{}{
			"user_id":    newToken.User_id,
			"token":      newToken.Token,
			"token_type": newToken.TokenType,
			"expires_at": newToken.Expires_at, // Converte time.Time para string no formato ISO 8601
			"created_at": newToken.Created_at,
			"active":     newToken.Active,
		}, false, "", "", "").
		ExecuteTo(&insertedTokens)

	if err != nil {
		log.Println("[InsertToken] Erro:", err)
		return nil, err
	}

	if len(insertedTokens) > 0 {
		return &insertedTokens[0], nil
	}
	return nil, nil
}

// RevokeToken desativa um token específico
func RevokeToken(token string) error {
	supabase := db.GetSupabase()

	_ ,_, err := supabase.From("tokens").
		Update(map[string]interface{}{"active": false}, "", "").
		Eq("token", token).
		Execute()

	if err != nil {
		log.Println("[RevokeToken] Erro:", err)
		return err
	}
	return nil
}

// RevokeAllTokensByUserID desativa todos os tokens de um usuário por tipo
func RevokeAllTokensByUserID(userID string, tokenType string) error {
	supabase := db.GetSupabase()

	_, _, err := supabase.From("tokens").
		Update(map[string]interface{}{"active": false}, "", "").
		Eq("user_id", userID).
		Eq("token_type", tokenType).
		Execute()

	if err != nil {
		log.Println("[RevokeAllTokensByUserID] Erro:", err)
		return err
	}
	return nil
}
