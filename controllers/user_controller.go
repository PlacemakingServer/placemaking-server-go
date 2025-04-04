package controllers

import (
    "net/http"
    "placemaking-backend-go/services"
	"placemaking-backend-go/models"
    "github.com/gin-gonic/gin"
)

func GetUserById(c *gin.Context) {

	//Pegando Id dos parâmetros da requisição
	id := c.Param("id")

	user, err := services.FetchUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao buscar usuário no banco de dados."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário encontrado.","user": models.SanitizeUser(user)})
}

func GetAllUsers(c *gin.Context) {

	users, err := services.FetchAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao buscar os usuários no banco de dados"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuários encontrados", "users": users})
}

func DeleteUserById(c *gin.Context) {

	id := c.Param("id")

	err := services.FetchDeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso."})
}

func UpdateUserById(c *gin.Context) {
	id := c.Param("id") // Pega o ID da URL

	var user models.User

	// Faz o binding do JSON para a struct User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	// Chama o service para atualizar o usuário
	updatedUser, err := services.FetchUpdateUser(id, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar usuário"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Usuário atualizado com sucesso", "user": updatedUser})
}