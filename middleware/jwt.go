package middleware

import (
	"net/http"
	"placemaking-backend-go/config"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var SECRET_KEY = config.LoadSettings().JwtSecret // Mude para uma chave segura!

// Middleware para validar JWT
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
			return []byte(SECRET_KEY), nil
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

		// Salva o ID do usuário no contexto
		if sub, ok := (*claims)["sub"].(string); ok {
			c.Set("user_id", sub)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token sem ID de usuário"})
			c.Abort()
			return
		}

		c.Next()
	}
}