package main

import (
	"log"
	"placemaking-backend-go/db"
	"placemaking-backend-go/routes"
)

func main() {
	// Inicializa Supabase
	db.InitSupabase()

	// Configura o servidor
	router := routes.SetupRouter() // Agora a função SetupRouter é reconhecida

	log.Println("Server running on http://localhost:8000...")
	if err := router.Run(":8000"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
