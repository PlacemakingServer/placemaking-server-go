package controllers

import (
	"net/http"
	"placemaking-backend-go/models"
	"placemaking-backend-go/services"

	"github.com/gin-gonic/gin"
)

// Criar uma nova pesquisa
func CreateSurvey(c *gin.Context) {
	var createSurveyData models.CreateSurvey

	if err := c.ShouldBindJSON(&createSurveyData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao processar os dados. Verifique as informações enviadas.",
			"error":   err.Error(),
		})
		return
	}

	survey, err := services.CreateSurvey(createSurveyData)
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

// Obter uma pesquisa por ID
func GetSurveyById(c *gin.Context) {
	id := c.Param("surveyId")
	researchId := c.Param("researchId")
	
	var surveyType models.SurveyType

	if err := c.ShouldBindJSON(&surveyType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao processar os dados. Verifique as informações enviadas.",
			"error":   err.Error(),
		})
		return
	}


	survey, err := services.GetSurveyById(id, researchId, surveyType.Type)
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

// Atualizar uma pesquisa por ID
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

	survey, err := services.UpdateSurveyById(id, updateData.SurveyType, updateData)
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

// Deletar uma pesquisa por ID
func DeleteSurveyById(c *gin.Context) {
	id := c.Param("surveyId")
	
		var surveyType models.SurveyType

	if err := c.ShouldBindJSON(&surveyType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao processar os dados. Verifique as informações enviadas.",
			"error":   err.Error(),
		})
		return
	}

	_, err := services.DeleteSurveyById(id, surveyType.Type)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao deletar a pesquisa.",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":        "Pesquisa deletada com sucesso.",
	})
}

// Obter pesquisas por research_id
func GetSurveysByResearchId(c *gin.Context) {
	researchId := c.Param("researchId")

	var surveyType models.SurveyType

	if err := c.ShouldBindJSON(&surveyType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao processar os dados. Verifique as informações enviadas.",
			"error":   err.Error(),
		})
		return
	}

	surveys, err := services.GetSurveysByResearchId(researchId, surveyType.Type)
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