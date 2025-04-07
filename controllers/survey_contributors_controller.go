package controllers

import (
	"net/http"
	"placemaking-backend-go/models"
	"placemaking-backend-go/services"

	"github.com/gin-gonic/gin"
)

// Criar um novo contribuidor para uma pesquisa
func CreateSurveyContributor(c *gin.Context) {
	var data models.CreateSurveyContributors

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Dados inválidos", "error": err.Error()})
		return
	}

	id := c.Param("surveyId") // ID vindo da URL

	surveyContributor, err := services.CreateSurveyContributorService(id, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao criar contribuidor", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Contribuidor criado com sucesso.","contributor": surveyContributor})
}

// Obter todos os contribuidores de uma pesquisa pelo surveyId
func GetSurveyContributorsBySurveyId(c *gin.Context) {
	surveyId := c.Param("surveyId")

	contributors, err := services.GetSurveyContributorsBySurveyIdService(surveyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao buscar contribuidores", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lista de todos os contribuidores obtida.", "contributors": contributors})
}

// Deletar um contribuidor de uma pesquisa pelo ID e surveyId (assíncrono)
func DeleteSurveyContributor(c *gin.Context) {
	id := c.Param("contributorId")
	surveyId := c.Param("surveyId")

	// Canal para receber o resultado da goroutine
	errChan := make(chan error, 1)

	// Executa a deleção em uma goroutine
	go func() {
		errChan <- services.DeleteSurveyContributorsByIdService(id, surveyId)
	}()

	// Aguarda a resposta da goroutine
	err := <-errChan
	close(errChan)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao deletar contribuidor", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contribuidor deletado com sucesso"})
}

