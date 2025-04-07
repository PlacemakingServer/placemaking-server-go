package controllers

import (
	"net/http"
	"placemaking-backend-go/models"
	"placemaking-backend-go/services"
	"github.com/gin-gonic/gin"
)

// CreateFieldOption cria uma nova opção para um campo
func CreateFieldOption(c *gin.Context) {
	fieldId := c.Param("fieldId")

	var createFieldOptionData models.CreateFieldOption
	if err := c.ShouldBindJSON(&createFieldOptionData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Dados inválidos", "error": err.Error()})
		return
	}

	fieldOption, err := services.CreateFieldOptionService(fieldId, createFieldOptionData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao criar opção de campo", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Opção criada com sucesso!", "field_option": fieldOption})
}

// GetAllFieldOptionsByFieldId retorna todas as opções de um campo
func GetAllFieldOptionsByFieldId(c *gin.Context) {
	fieldId := c.Param("fieldId")

	fieldOptions, err := services.GetAllFieldOptionsByFieldIdService(fieldId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao buscar opções de campo", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lista de opções do campo.", "field_options": fieldOptions})
}

// DeleteFieldOptionById deleta uma opção de campo pelo ID
func DeleteFieldOptionById(c *gin.Context) {
	id := c.Param("optionId")
	fieldId := c.Param("fieldId")

	// Criar um canal para capturar erros da goroutine
	errChan := make(chan error, 1)

	// Iniciar uma goroutine para deletar a opção de campo
	go func() {
		errChan <- services.DeleteFieldOptionByIdService(id, fieldId)
		close(errChan) // Fecha o canal após enviar o erro
	}()

	// Aguardar o resultado da goroutine
	if err := <-errChan; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao deletar a opção de campo.",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Opção de campo deletada com sucesso."})
}
