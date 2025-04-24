package controllers

import (
	"log"
	"net/http"
	"net/url"
	"placemaking-backend-go/models"
	"placemaking-backend-go/services"

	"github.com/gin-gonic/gin"
)

func CreateSurvey(c *gin.Context) {
	var createSurveyData models.CreateSurvey

	if err := c.ShouldBindJSON(&createSurveyData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao processar os dados. Verifique as informações enviadas.",
			"error":   err.Error(),
		})
		return
	}

	surveyType, _ := url.QueryUnescape(c.Query("survey_type"))

	survey, err := services.CreateSurvey(surveyType, createSurveyData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao criar a pesquisa.",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Pesquisa criada com sucesso.",
		"survey":  survey,
	})
}

func GetSurveyById(c *gin.Context) {
	id := c.Param("surveyId")
	researchId := c.Param("researchId")

	surveyType, _ := url.QueryUnescape(c.Query("survey_type"))

	survey, err := services.GetSurveyById(id, researchId, surveyType)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Pesquisa não encontrada.",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Pesquisa recuperada com sucesso.",
		"survey":  survey,
	})
}

func UpdateSurveyById(c *gin.Context) {
	id := c.Param("surveyId")

	var updateData models.UpdateSurvey
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao processar os dados. Verifique as informações enviadas.",
			"error":   err.Error(),
		})
		return
	}

	surveyType, _ := url.QueryUnescape(c.Query("survey_type"))

	survey, err := services.UpdateSurveyById(id, surveyType, updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao atualizar a pesquisa.",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Pesquisa atualizada com sucesso.",
		"survey":  survey,
	})
}

func DeleteSurveyById(c *gin.Context) {
	id := c.Param("surveyId")

	surveyType, _ := url.QueryUnescape(c.Query("survey_type"))

	_, err := services.DeleteSurveyById(id, surveyType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao deletar a pesquisa.",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Pesquisa deletada com sucesso.",
	})
}

func GetSurveysByResearchId(c *gin.Context) {
	researchId := c.Param("researchId")
	surveyType, _ := url.QueryUnescape(c.Query("survey_type"))

	log.Println(surveyType)

	surveys, err := services.GetSurveysByResearchId(researchId, surveyType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao buscar pesquisas pelo research_id.",
			"error":   err.Error(),
		})
		return
	}

	if len(surveys) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Nenhuma pesquisa encontrada para o research_id informado.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Pesquisas recuperadas com sucesso.",
		"surveys": surveys,
	})
}
