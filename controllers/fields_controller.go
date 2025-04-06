package controllers

import (
	"net/http"
	"placemaking-backend-go/models"
	"placemaking-backend-go/services"

	"github.com/gin-gonic/gin"
)

// CreateField cria um novo campo em um survey
func CreateField(c *gin.Context) {
	surveyId := c.Param("surveyId")
	surveyType := c.Query("survey_type")

	var createFieldData models.CreateField
	if err := c.ShouldBindJSON(&createFieldData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Dados inválidos", "error": err.Error()})
		return
	}

	field, err := services.CreateFieldService(surveyId, surveyType, createFieldData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao criar campo", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Campo criado com sucesso", "field": field})
}

// GetAllFieldsBySurveyId retorna todos os campos de um survey específico
func GetAllFieldsBySurveyId(c *gin.Context) {
	surveyId := c.Param("surveyId")
	surveyType := c.Query("survey_type")

	fields, err := services.GetAllFieldsBySurveyIdService(surveyId, surveyType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao buscar campos", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Campos encontrados", "fields": fields})
}

// DeleteField deleta um campo de um survey
func DeleteField(c *gin.Context) {
	fieldId := c.Param("id")
	surveyId := c.Param("surveyId")
	surveyType := c.Query("survey_type")

	err := services.DeleteFieldBySurveyIdService(fieldId, surveyId, surveyType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao deletar campo", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Campo deletado com sucesso"})
}

// UpdateField atualiza um campo específico de um survey
func UpdateField(c *gin.Context) {
	fieldId := c.Param("id")
	surveyId := c.Param("surveyId")
	surveyType := c.Query("survey_type")

	var updateFieldData models.CreateField
	if err := c.ShouldBindJSON(&updateFieldData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Dados inválidos", "error": err.Error()})
		return
	}

	field, err := services.UpdateFieldService(fieldId, surveyId, surveyType, updateFieldData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao atualizar campo", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Campo atualizado com sucesso", "field": field})
}
