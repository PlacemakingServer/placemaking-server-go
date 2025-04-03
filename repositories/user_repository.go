package repository

import (
	"log"
	"placemaking-backend-go/config"
	"placemaking-backend-go/models"
)

func GetUserById(id string) (models.User, error) {
	supabase := config.GetSupabase()

	var user models.User

	_, err := supabase.From("users").
		Select("*", "", false).
		Eq("id", id).
		Single().
		ExecuteTo(&user) // Decodifica para a struct User

	if err != nil {
		log.Println("Erro ao buscar usuário:", err)
		return models.User{}, err
	}

	// Converte timestamps
	user.ConvertTimestamps()

	return user, nil
}
