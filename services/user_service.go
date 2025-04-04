package services

import (
	"log"
	"math/rand"
	"placemaking-backend-go/models"
	repository "placemaking-backend-go/repositories"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func FetchUserById(id string) (models.User, error) {
	user, err := repository.GetUserById(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func FetchAllUsers() ([]models.SanitizedUser, error) {
	users, err := repository.GetAllUsers()
	if err != nil {
		return []models.SanitizedUser{}, err
	}

	return users, nil
}

func FetchDeleteUser(id string) error {
	err := repository.DeleteUserById(id)
	if err != nil {
		return err
	}
	return nil
}

func FetchUpdateUser(id string, user models.User) (models.User, error) {

	updatedData := map[string]interface{}{
		"name":       user.Name,
		"email":      user.Email,
		"role":       user.Role,
		"updated_at": time.Now().Format(time.RFC3339),
		"status":     user.Status,
	}

	user, err := repository.UpdateUserById(id, updatedData)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GenerateUserPassword() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	password := make([]byte, 8)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	return string(password)
}

func GenerateHashedPassword(password string) (string, error) {
	salt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(salt), nil
}

func CheckPassword(plainPassword string, hashedPassword string) bool {
	// Verifica se os valores não estão vazios
	if plainPassword == "" || hashedPassword == "" {
		return false
	}

	// Compara a senha em texto com o hash armazenado
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		log.Println("Erro ao verificar senha:", err)
		return false
	}

	return true
}
