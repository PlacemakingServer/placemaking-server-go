package controllers

import (
	"log"
	"net/http"
	"strings"
	"placemaking-backend-go/models"
	"placemaking-backend-go/services"

	"github.com/gin-gonic/gin"
)

func Register(c* gin.Context) {

	var createUserData models.CreateUser

	if err := c.ShouldBindJSON(&createUserData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	user, temporaryPassword, err:= services.RegisterUser(createUserData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	data := map[string]interface{}{
		"user": map[string]interface{}{
			"email": user.Email,
			"name":  user.Name,
		},
		"temporary_password": temporaryPassword,
	}
	
	err = services.SendUserData(data)
	if err != nil {
		log.Println("Erro ao enviar e-mail:", err)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Usuário criado com sucesso.", "user": user})
}

func Login(c* gin.Context) {

	var loginData models.Login

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	data, err := services.LoginUser(loginData.Email, loginData.Password)

	if err != nil {
		c.JSON(400, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuário logado com sucesso!", "acess_token": data["token"], "user": data["user"]})

}

func Logout(c *gin.Context) {
	// Obtém o token do cabeçalho Authorization
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Token não fornecido"})
		return
	}

	// Extrai o token do cabeçalho
	token := strings.Split(authHeader, " ")[1]

	// Processa o logout em background
	go services.LogoutUser(token)

	c.JSON(http.StatusOK, gin.H{"message": "Logout realizado com sucesso!"})
}