package controllers

import (
	"net/http"
	"placemaking-backend-go/models"
	"placemaking-backend-go/services"

	"github.com/gin-gonic/gin"
)

// Obter todas as respostas de uma pesquisa
func GetSurveyAnswers(c *gin.Context) {
	surveyId := c.Param("surveyId")

	answers, err := services.GetSurveyAnswers(surveyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao buscar respostas.",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Respostas recuperadas com sucesso.",
		"answers": answers,
	})
}

// Criar uma nova resposta
func CreateSurveyAnswer(c *gin.Context) {
	surveyId := c.Param("surveyId")
	surveyType := c.Query("surveyType")
	contributorId := c.Query("contributorId")

	var answerData models.CreateSurveyAnswer
	if err := c.ShouldBindJSON(&answerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Dados inválidos.",
			"error":   err.Error(),
		})
		return
	}

	answer, err := services.CreateSurveyAnswer(surveyId, surveyType, contributorId, answerData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao criar resposta.",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Resposta criada com sucesso.",
		"answer":  answer,
	})
}

// Deletar resposta
func DeleteSurveyAnswer(c *gin.Context) {
	answerId := c.Param("id")

	err := services.DeleteSurveyAnswer(answerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao deletar resposta.",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Resposta deletada com sucesso.",
	})
}
