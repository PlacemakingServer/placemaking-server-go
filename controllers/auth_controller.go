package controllers

import (
	"fmt"
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

	c.JSON(http.StatusOK, gin.H{"message": "Usuário logado com sucesso!", "access_token": data["token"], "user": data["user"]})

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

func ForgotPassword(c *gin.Context) {

	var email models.AuthEmailRecovery

	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	forgotPasswordData, err := services.ForgotPasswordService(email.Email) 

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if forgotPasswordData["error"] != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": forgotPasswordData["error"]})
		return
	}

	err = services.SendForgotEmailPasswordData(forgotPasswordData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Código enviado com sucesso!"})
}

func ValidateCode(c *gin.Context) {
	var token models.AuthValidadetToken
	var validateToken string

	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	validateCodeData, err := services.ValidateCodeService(token.Token)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Debug para entender a estrutura do validateCodeData
	fmt.Println("validateCodeData:", validateCodeData)

	if validateCodeData["error"] != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validateCodeData["error"]})
		return
	}

	// Corrigindo a chave do token
	if v, ok := validateCodeData["validated_token"].(string); ok {
		validateToken = v
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Token inválido"})
		return
	}

	err = services.RevokeToken(validateToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao revogar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Código validado com sucesso!",
		"access_token": validateCodeData["access_token"],
	})
}

func ResetPassword(c *gin.Context) {
	var userRecoveryData models.AuthUserResetPassword

	// Bind JSON para a struct
	if err := c.ShouldBindJSON(&userRecoveryData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	// Obter user_id do contexto
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	// Converter para string
	userID, ok := userIDInterface.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao recuperar ID do usuário"})
		return
	}

	// Chamar serviço para redefinir senha
	resetPasswordData, err := services.ResetPasswordService(userRecoveryData.NewPassword, userRecoveryData.ConfirmPassword, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao redefinir senha"})
		return
	}

	// Verifica se houve erro na resposta do serviço
	if errMsg, ok := resetPasswordData["error"].(string); ok && errMsg != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
		return
	}

	// Resposta de sucesso
	c.JSON(http.StatusOK, gin.H{"message": "Senha atualizada com sucesso!"})
}
