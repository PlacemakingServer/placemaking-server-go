package models

import (
	"log"
	"placemaking-backend-go/utils"
)

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
