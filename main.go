package handler

import (
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"placemaking-backend-go/config"
	"placemaking-backend-go/controllers"
)

// setupRouter configura as rotas da API
func setupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/input_types", controllers.GetInputTypes)
	}

	return router
}

// Handler é a função usada pelo Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	router := setupRouter()
	router.ServeHTTP(w, r)
}

// main é usado para rodar localmente
func main() {
	// Carregar variáveis do .env localmente
	_ = godotenv.Load()

	// Inicializar Supabase
	config.InitSupabase()

	// Verificar se está rodando na Vercel
	if os.Getenv("VERCEL") != "" {
		log.Println("Rodando na Vercel...")
		return
	}

	// Rodar localmente
	router := setupRouter()
	log.Println("Server running on port 8080...")
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}