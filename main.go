package handler

import (
	"net/http"
	"placemaking-backend-go/db"
	"placemaking-backend-go/routes"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	db.InitSupabase()

	router := routes.SetupRouter()
	router.ServeHTTP(w, r)
}
