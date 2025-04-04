package controllers

import (
	"net/http"
	"placemaking-backend-go/models"
	"placemaking-backend-go/services"

	"github.com/gin-gonic/gin"
)

// Criar uma nova pesquisa
func CreateResearch(c *gin.Context) {
	var createResearchData models.CreateResearch

	if err := c.ShouldBindJSON(&createResearchData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	research, err := services.FetchCreateResearch(createResearchData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar pesquisa"})
		return
	}

	c.JSON(http.StatusCreated, research)
}

// Buscar todas as pesquisas
func GetAllResearches(c *gin.Context) {
	researches, err := services.FetchAllResearches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar pesquisas"})
		return
	}

	c.JSON(http.StatusOK, researches)
}

// Buscar pesquisa por ID
func GetResearchById(c *gin.Context) {
	id := c.Param("id")

	research, err := services.FetchResearchById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pesquisa não encontrada"})
		return
	}

	c.JSON(http.StatusOK, research)
}

// Atualizar pesquisa por ID
func UpdateResearch(c *gin.Context) {
	id := c.Param("id")
	var updateResearchData models.UpdateResearch

	if err := c.ShouldBindJSON(&updateResearchData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	research, err := services.FetchUpdateResearch(id, updateResearchData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar pesquisa"})
		return
	}

	c.JSON(http.StatusOK, research)
}

// Deletar pesquisa por ID
func DeleteResearch(c *gin.Context) {
	id := c.Param("id")

	err := services.FetchDeleteResearch(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar pesquisa"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
