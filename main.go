package main

import (
	"log"
	"placemaking-backend-go/config"
	"placemaking-backend-go/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/input_types", controllers.GetInputTypes)
	}
}


func main() {
	// Inicializar Supabase
	config.InitSupabase()

	// Criar uma nova instância do Gin
	router := gin.Default()

	// Registrar rotas
	SetupRoutes(router)

	// Rodar o servidor na porta 8080
	log.Println("Server running on port 8080...")
	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}