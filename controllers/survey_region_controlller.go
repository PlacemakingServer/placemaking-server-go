package controllers

import (
	"net/http"
	"net/url"
	"placemaking-backend-go/models"
	"placemaking-backend-go/services"

	"github.com/gin-gonic/gin"
)

// POST /surveys/:surveyId/regions?survey_type=tipo
func CreateSurveyRegion(c *gin.Context) {
	surveyId := c.Param("surveyId")
	surveyType, _ := url.QueryUnescape(c.Query("survey_type"))

	var regionData models.CreateSurveyRegion
	if err := c.ShouldBindJSON(&regionData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	region, err := services.CreateSurveyRegion(surveyId, surveyType, regionData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar região"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Região de pesquisa criada com sucesso!",
		"region":  region,
	})
}

// GET /surveys/:surveyId/regions
func GetAllSurveyRegionsBySurveyId(c *gin.Context) {
	surveyId := c.Param("surveyId")

	regions, err := services.GetAllSurveyRegionsBySurveyId(surveyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar regiões"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Regiões encontradas com sucesso!",
		"regions": regions,
	})
}

// GET /regions?survey_type=tipo
func GetAllSurveysBySurveyType(c *gin.Context) {
	surveyType, _ := url.QueryUnescape(c.Query("survey_type"))

	regions, err := services.GetAllSurveysBySurveyType(surveyType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar regiões por tipo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Regiões por tipo encontradas com sucesso!",
		"regions": regions,
	})
}

// GET /regions/:regionId
func GetSurveyRegionById(c *gin.Context) {
	id := c.Param("regionId")

	region, err := services.GetSurveyRegionById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar região"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Região encontrada com sucesso!",
		"region":  region,
	})
}

// PUT /regions/:regionId
func UpdateSurveyRegionById(c *gin.Context) {
	id := c.Param("regionId")

	var data models.UpdateSurveyRegion
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	updatedRegion, err := services.UpdateSurveyRegion(id, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar região"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Região atualizada com sucesso!",
		"region":  updatedRegion,
	})
}

// DELETE /regions/:regionId
func DeleteSurveyRegion(c *gin.Context) {
	id := c.Param("regionId")

	deleted, err := services.DeleteSurveyRegion(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar região"})
		return
	}

	if len(deleted) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Nenhuma região encontrada para deletar.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Região deletada com sucesso!",
	})
}
