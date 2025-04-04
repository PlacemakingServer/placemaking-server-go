package handler

import (
	"net/http"
	"placemaking-backend-go/db"
	"placemaking-backend-go/routes"
)

// Handler é a única função exportada, usada pelo Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	// Inicializa Supabase antes de rodar as rotas
	db.InitSupabase()

	// Cria o router e serve a requisição
	router := routes.SetupRouter()
	router.ServeHTTP(w, r)
}

