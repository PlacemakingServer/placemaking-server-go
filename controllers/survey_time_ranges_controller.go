package controllers

import (
	"net/http"
	"net/url"
	"placemaking-backend-go/models"
	"placemaking-backend-go/services"

	"github.com/gin-gonic/gin"
)

// POST /surveys/:surveyId/time-ranges?survey_type=tipo
func CreateSurveyTimeRange(c *gin.Context) {
	surveyId := c.Param("surveyId")
	surveyType, _ := url.QueryUnescape(c.Query("survey_type"))

	if surveyType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "survey_type é obrigatório"})
		return
	}

	var data models.CreateSurveyTimeRange
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	timeRange, err := services.CreateSurveyTimeRange(surveyId, surveyType, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar intervalo de tempo"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Intervalo de tempo criado com sucesso!",
		"survey_time": timeRange,
	})
}

// GET /surveys/:surveyId/time-ranges
func GetAllSurveyTimeRangeBySurveyId(c *gin.Context) {
	surveyId := c.Param("surveyId")

	timeRanges, err := services.GetAllSurveyTimeRangeBySurveyId(surveyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar intervalos de tempo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Intervalos de tempo encontrados com sucesso!",
		"survey_times": timeRanges,
	})
}

// GET /time-ranges?survey_type=tipo
func GetAllSurveyTimeRangeBySurveyType(c *gin.Context) {
	surveyType, _ := url.QueryUnescape(c.Query("survey_type"))

	timeRanges, err := services.GetAllSurveyTimeRangeBySurveyType(surveyType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar intervalos de tempo por tipo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Intervalos de tempo por tipo encontrados com sucesso!",
		"survey_times": timeRanges,
	})
}

// GET /time-ranges/:timeRangeId
func GetSurveyTimeRangeById(c *gin.Context) {
	id := c.Param("timeRangeId")

	timeRange, err := services.GetSurveyTimeRangeById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar intervalo de tempo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Intervalo de tempo encontrado com sucesso!",
		"survey_time": timeRange,
	})
}

// PUT /time-ranges/:timeRangeId
func UpdateSurveyTimeRangeById(c *gin.Context) {
	id := c.Param("timeRangeId")

	var data models.UpdateSurveyTimeRange
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	updated, err := services.UpdateSurveyTimeRange(id, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar intervalo de tempo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Intervalo de tempo atualizado com sucesso!",
		"survey_time": updated,
	})
}

// DELETE /time-ranges/:timeRangeId
func DeleteSurveyTimeRange(c *gin.Context) {
	id := c.Param("timeRangeId")

	deleted, err := services.DeleteSurveyTimeRange(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar intervalo de tempo"})
		return
	}

	if len(deleted) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Intervalo de tempo não encontrado.",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message":     "Intervalo de tempo deletado com sucesso!",
	})
}
