package middleware

import (
	"log"
	"net/http"
	"placemaking-backend-go/config"
	"placemaking-backend-go/db"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTAuthMiddleware valida o JWT e verifica se o token está ativo
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		claims := &jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.LoadSettings().JwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// Verifica se o token expirou
		exp, ok := (*claims)["exp"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato do token inválido"})
			c.Abort()
			return
		}

		expirationTime := time.Unix(int64(exp), 0)
		if expirationTime.Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expirado"})
			c.Abort()
			return
		}

		// Obtém o ID do usuário do token
		var userID string
		if sub, ok := (*claims)["sub"].(string); ok {
			userID = sub
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token sem ID de usuário"})
			c.Abort()
			return
		}

		// Verifica se o token está ativo no Supabase
		supabase := db.GetSupabase()
		var tokens []map[string]interface{}

		_, err = supabase.From("tokens").
			Select("active", "", false).
			Eq("token", tokenString).
			ExecuteTo(&tokens)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Erro ao validar token"})
			c.Abort()
			return
		}

		// Se o token estiver inativo, rejeita a requisição
		if len(tokens) == 0 || tokens[0]["active"] == nil || tokens[0]["active"].(bool) == false {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token revogado"})
			c.Abort()
			return
		}

		// Salva o ID do usuário no contexto
		c.Set("user_id", userID)
		log.Println("✅ user_id salvo no contexto:", userID)
		c.Next()
	}
}
