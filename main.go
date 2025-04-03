package handler

import (
	"fmt"
	"net/http"

	"placemaking-backend-go/config"
	"placemaking-backend-go/controllers"

	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Inicializar Supabase
	config.InitSupabase()

	// Criar uma nova instância do Gin
	router := gin.Default()

	// Criar uma rota de teste para ver se a Vercel está servindo corretamente
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is running on Vercel!")
	})

	// Adicionar rota manualmente
	router.GET("/api/input_types", controllers.GetInputTypes)

	// Log para debugar
	fmt.Println("Handler executado - Rotas registradas")

	// Servir a requisição
	router.ServeHTTP(w, r)
}
