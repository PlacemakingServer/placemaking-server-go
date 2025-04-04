package services

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"placemaking-backend-go/config"
	"placemaking-backend-go/models"
	repository "placemaking-backend-go/repositories"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const JWTAlgorithm = "HS256"
const JWTExpiresMinutes = 60 * 24 * 7
const RecoverPasswordExpiresMinutes = 5

func GenerateUserToken(user models.User, tokenType string) (*models.Token, error) {
	if tokenType == "Bearer" {
		existingToken, err := repository.GetTokenByUserID(user.ID, tokenType)
		if err == nil && existingToken != nil && existingToken.Active {
			// Converter string Expires_at para time.Time
			layout := "2006-01-02T15:04:05"
			expiresAt, err := time.Parse(layout, existingToken.Expires_at)
			if err != nil {
				log.Println("[GenerateUserToken] Erro ao converter Expires_at:", err)
				return nil, err
			}

			// Verifica se o token ainda é válido
			if expiresAt.After(time.Now()) {
				return existingToken, nil
			}
		}
	}

	_ = repository.RevokeAllTokensByUserID(user.ID, tokenType)

	var newTokenData map[string]interface{}
	if tokenType == "Bearer" {
		newTokenData = GenerateJWTToken(user)
	} else {
		newTokenData = GenerateRecoverPasswordToken()
	}

	// Certificar que newTokenData["expires_at"] está no formato correto
	var expiresAt time.Time
	switch v := newTokenData["expires_at"].(type) {
	case string:
		layout := time.RFC3339
		var err error
		expiresAt, err = time.Parse(layout, v)
		if err != nil {
			log.Println("[GenerateUserToken] Erro ao converter expires_at:", err)
			return nil, err
		}
	case time.Time:
		expiresAt = v
	default:
		return nil, fmt.Errorf("[GenerateUserToken] Tipo inesperado para expires_at: %T", v)
	}

	token, err := repository.InsertToken(
		user.ID,
		newTokenData["token"].(string),
		tokenType,
		expiresAt,
	)

	if err != nil {
		log.Println("[GenerateUserToken] Erro ao inserir token:", err)
		return nil, err
	}

	return token, nil
}


func GenerateJWTToken(user models.User) map[string]interface{} {
	expire := time.Now().Add(time.Minute * time.Duration(JWTExpiresMinutes))

	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   expire.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(config.LoadSettings().JwtSecret))

	return map[string]interface{}{
		"token":      signedToken,
		"expires_at": expire,
	}
}

func GenerateRecoverPasswordToken() map[string]interface{} {
	expire := time.Now().Add(time.Minute * time.Duration(RecoverPasswordExpiresMinutes))

	token := generateRandomToken(8)

	return map[string]interface{}{
		"token":      token,
		"expires_at": expire,
	}
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de assinatura inválido")
		}
		return []byte(config.LoadSettings().JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token inválido")
}

func RevokeToken(token string) error {
	return repository.RevokeToken(token)
}

func generateRandomToken(length int) string {
	rand.Seed(time.Now().UnixNano()) // Garante números aleatórios diferentes a cada execução
	const digits = "0123456789"
	token := make([]byte, length)
	for i := range token {
		token[i] = digits[rand.Intn(len(digits))]
	}
	return string(token)
}