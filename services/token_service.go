package services

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"placemaking-backend-go/config"
	"placemaking-backend-go/models"
	repository "placemaking-backend-go/repositories"
)

const JWTAlgorithm = "HS256"
const JWTExpiresMinutes = 60 * 24 * 7
const RecoverPasswordExpiresMinutes = 5

func GenerateUserToken(user models.User, tokenType string) (*models.Token, error) {
	if tokenType == "Bearer" {
		existingToken, err := repository.GetTokenByUserID(user.ID, tokenType)
		if err == nil && existingToken != nil && existingToken.Active {
			if existingToken.Expires_at.After(time.Now()) {
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

	token, err := repository.InsertToken(
		user.ID,
		newTokenData["token"].(string),
		tokenType,
		newTokenData["expires_at"].(time.Time),
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
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	token := make([]byte, length)
	for i := range token {
		token[i] = charset[randomInt(0, len(charset))]
	}
	return string(token)
}

func randomInt(min, max int) int {
	return min + (max-min)*randomSeed()/randomSeedMax()
}

func randomSeed() int {
	return int(time.Now().UnixNano() % 10000)
}

func randomSeedMax() int {
	return 10000
}
