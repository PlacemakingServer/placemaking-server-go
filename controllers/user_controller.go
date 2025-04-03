package controllers

import (
    "net/http"
    "placemaking-backend-go/services"

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

	c.JSON(http.StatusOK, gin.H{"message": "Usuário encontrado.","user": user})
}