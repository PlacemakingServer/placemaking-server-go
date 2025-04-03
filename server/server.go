package main

import (
	"log"
	"placemaking-backend-go/config"
	"placemaking-backend-go/routes"
)

func main() {
	// Inicializa Supabase
	config.InitSupabase()

	// Configura o servidor
	router := routes.SetupRouter() // Agora a função SetupRouter é reconhecida

	log.Println("Server running on http://localhost:8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
