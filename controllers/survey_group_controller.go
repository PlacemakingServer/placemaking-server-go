package controllers

import (
	"net/http"
	"net/url"
	"placemaking-backend-go/models"
	"placemaking-backend-go/services"

	"github.com/gin-gonic/gin"
)

// Obter todos os grupos de uma pesquisa
func GetSurveyGroups(c *gin.Context) {
	surveyId := c.Param("surveyId")
	surveyType, _ := url.QueryUnescape(c.Query("survey_type"))

	groups, err := services.GetSurveyGroups(surveyId, surveyType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao buscar grupos.",
			"error":   err.Error(),
		})
		return
	}

	if len(groups) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Nenhum grupo encontrado.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Grupos recuperados com sucesso.",
		"groups":  groups,
	})
}

// Criar um novo grupo de pesquisa
func CreateSurveyGroup(c *gin.Context) {
	var groupData models.CreateSurveyGroup
	if err := c.ShouldBindJSON(&groupData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Dados inválidos.",
			"error":   err.Error(),
		})
		return
	}

	surveyId := c.Param("surveyId")

	group, err := services.CreateSurveyGroup(surveyId, groupData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Erro ao criar grupo.",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Grupo criado com sucesso.",
		"group":   group,
	})
}

// Deletar grupo de pesquisa
func DeleteSurveyGroup(c *gin.Context) {
	id := c.Param("groupId")

	err := services.DeleteSurveyGroup(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao deletar grupo.",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Grupo deletado com sucesso.",
	})
}
