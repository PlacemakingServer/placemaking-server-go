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

c.JSON(http.StatusCreated, gin.H{"message": "Pesquisa criada com sucesso!", "research": research})
}

// Buscar todas as pesquisas
func GetAllResearches(c *gin.Context) {
	researches, err := services.FetchAllResearches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar pesquisas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lista de pesquisas encontradas.", "researches": researches})
}

// Buscar pesquisa por ID
func GetResearchById(c *gin.Context) {
	id := c.Param("researchId")

	research, err := services.FetchResearchById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pesquisa não encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pesquisa encontrada com sucesso.", "research": research})
}

// Atualizar pesquisa por ID
func UpdateResearch(c *gin.Context) {
	id := c.Param("researchId")
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

	c.JSON(http.StatusOK, gin.H{"message": "Pesquisa atualizada com sucesso.", "research": research})
}

// Deletar pesquisa por ID
func DeleteResearch(c *gin.Context) {
	id := c.Param("researchId")

	errChan := make(chan error, 1)

	go func() {
		errChan <- services.FetchDeleteResearch(id)
	}()

	err := <-errChan
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar pesquisa"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pesquisa apagada com sucesso."})
}