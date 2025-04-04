package models

import (
	"log"
	"placemaking-backend-go/utils"
)

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Status     string `json:"status"`
}

type CreateUser struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	ConfirmEmail string `json:"confirmation_email"`
	Role         string `json:"role"`
	Status       string `json:"status"`
}

type InsertUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Status   string `json:"status"`
}

// SanitizedUser remove a senha antes de retornar os dados do usuário.
type SanitizedUser struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Status    string `json:"status"`
}

// SanitizeUser cria um novo objeto sem a senha.
func SanitizeUser(user User) SanitizedUser {
	return SanitizedUser{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.Created_at,
		UpdatedAt: user.Updated_at,
		Status:    user.Status,
	}
}

// Método para converter as strings em time.Time
func (u *User) ConvertTimestamps() {
	var err error
	u.Created_at, err = utils.ParseTimestamp(u.Created_at)
	if err != nil {
		log.Println("Erro ao converter CreatedAt:", err)
	}

	u.Updated_at, err = utils.ParseTimestamp(u.Updated_at)
	if err != nil {
		log.Println("Erro ao converter UpdatedAt:", err)
	}
}
