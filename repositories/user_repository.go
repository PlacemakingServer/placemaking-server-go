package repository

import (
	"log"
	"placemaking-backend-go/db"
	"placemaking-backend-go/models"
)

func GetUserById(id string) (models.User, error) {
	supabase := db.GetSupabase()

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

func GetUserByEmail(email string) (models.User, error) {
	supabase := db.GetSupabase()

	var user models.User

	_, err := supabase.From("users").
		Select("*", "", false).
		Eq("email", email).
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

func GetAllUsers() ([]models.User, error){
	supabase := db.GetSupabase()

	var users []models.User

	_, err := supabase.From("users").
	Select("*", "", false).
	ExecuteTo(&users) // Decodifica para a struct User

	if err != nil {
		log.Println("Erro ao buscar usuários:", err) 
		return nil, err
	}

	for i := range users {
		users[i].ConvertTimestamps()
	}

	return users, nil

}

func DeleteUserById(id string) error {
	supabase := db.GetSupabase()

	_, _, err := supabase.From("users").
		Delete("", "").
		Eq("id", id). // Filtra pelo ID do usuário
		Execute()


	if err != nil {
		log.Println("Erro ao deletar usuário:", err)
		return err
	}

	return nil
}

func UpdateUserById(id string, updatedData map[string]interface{}) (models.User, error) {
	supabase := db.GetSupabase()

	var updatedUser models.User

	_, err := supabase.From("users").
		Update(updatedData, "", ""). 
		Eq("id", id).
		Single().
		ExecuteTo(&updatedUser)

	if err != nil {
		log.Println("Erro ao atualizar usuário:", err)
		return models.User{}, err
	}

	return updatedUser, nil
}

func InsertUser(createUserData models.InsertUser) (*models.User, error) {
	supabase := db.GetSupabase()

	var user models.User

	role := string([]rune(createUserData.Role)[:])

	// Inserção no banco de dados
	_, err := supabase.From("users").Insert(map[string]interface{}{
		"name":     createUserData.Name,
		"email":    createUserData.Email,
		"password": createUserData.Password,
		"role":     role,
		"status":   createUserData.Status,
	},false,"","","").Single().ExecuteTo(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}


func UpdateUserPassword(userID string, password string) error {
	supabase := db.GetSupabase()

	_, _, err := supabase.From("users").
		Update(map[string]interface{}{"password": password}, "", "").
		Eq("id", userID).
		Execute()

	if err != nil {
		return err
	}

	return nil
}