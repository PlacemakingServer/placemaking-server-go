package controllers

import (
	"log"
	"net/http"
	"net/url"
	"placemaking-backend-go/models"
	"placemaking-backend-go/services"

	"github.com/gin-gonic/gin"
)

// POST /surveys/:surveyId/answers
func CreateSurveyAnswer(c *gin.Context) {
	surveyId := c.Param("surveyId")
	surveyType, _ := url.QueryUnescape(c.Query("survey_type"))
	contributorId, _ := url.QueryUnescape(c.Query("contributor_id"))
	
	var data models.CreateSurveyAnswer

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	answer, err := services.CreateSurveyAnswer(surveyId, surveyType, contributorId, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar resposta"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Resposta criada com sucesso!",
		"answer":  answer,
	})
}

// GET /surveys/:surveyId/answers
func GetSurveyAnswersBySurveyId(c *gin.Context) {
	surveyId := c.Param("surveyId")

	answers, err := services.GetSurveyAnswersBySurveyId(surveyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar respostas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Respostas buscadas com sucesso!",
		"answers": answers,
	})
}

// GET /contributors/:contributorId/answers
func GetSurveyAnswersByContributorId(c *gin.Context) {
	contributorId := c.Param("contributorId")

	answers, err := services.GetSurveyAnswersByContributorId(contributorId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar respostas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Respostas do colaborador buscadas com sucesso!",
		"answers": answers,
	})
}

// GET /answers/:id
func GetSurveyAnswerById(c *gin.Context) {
	id := c.Param("answerId")

	answer, err := services.GetSurveyAnswerById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar resposta"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Resposta buscada com sucesso!",
		"answer":  answer,
	})
}

// PUT /answers/:id
func UpdateSurveyAnswerById(c *gin.Context) {
	id := c.Param("answerId")

	var data models.UpdateSurveyAnswer
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	answer, err := services.UpdateSurveyAnswerById(id, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar resposta"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Resposta atualizada com sucesso!",
		"answer":  answer,
	})
}

// DELETE /answers/:id
func DeleteSurveyAnswerById(c *gin.Context) {
	id := c.Param("answerId")

	answer, err := services.DeleteSurveyAnswerById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar resposta"})
		return
	}

	if len(answer) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Nenhuma resposta encontrada.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Resposta deletada com sucesso!",
	})
}
