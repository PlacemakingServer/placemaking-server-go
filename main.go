package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"placemaking-backend-go/config"
	"placemaking-backend-go/controllers"
)

// setupRouter configura as rotas
func setupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/input_types", controllers.GetInputTypes)
	}

	return router
}

// Handler é a única função exportada, usada pelo Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	// Inicializa Supabase antes de rodar as rotas
	config.InitSupabase()

	// Cria o router e serve a requisição
	router := setupRouter()
	router.ServeHTTP(w, r)
}

