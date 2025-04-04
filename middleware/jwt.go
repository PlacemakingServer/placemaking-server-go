package middleware

import (
	"fmt"
	"net/http"
	"placemaking-backend-go/config"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTAuthMiddleware valida o JWT e verifica se o token está ativo
func JWTAuthMiddleware(publicRoutes map[string]bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		path := c.FullPath()
		routeKey := fmt.Sprintf("%s %s", method, path)

		// Se a rota for pública, permite a requisição sem autenticação
		if publicRoutes[routeKey] {
			c.Next()
			return
		}

		// Caso contrário, exige autenticação
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

		// Salva o ID do usuário no contexto
		c.Set("user_id", userID)
		c.Next()
	}
}
