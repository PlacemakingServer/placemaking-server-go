package controllers

import (
	"net/http"
	"placemaking-backend-go/models"
	"placemaking-backend-go/services"

	"github.com/gin-gonic/gin"
)

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

func GetAllResearches(c *gin.Context) {
	researches, err := services.FetchAllResearches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar pesquisas"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lista de pesquisas encontradas.", "researches": researches})
}

func GetResearchById(c *gin.Context) {
	id := c.Param("researchId")

	research, err := services.FetchResearchById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pesquisa não encontrada"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pesquisa encontrada com sucesso.", "research": research})
}

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

func DeleteResearch(c *gin.Context) {
	id := c.Param("researchId")

	resultChan := make(chan []models.Research, 1)
	errChan := make(chan error, 1)

	go func() {
		deletedResearch, err := services.FetchDeleteResearch(id)
		if err != nil {
			errChan <- err
			return
		}
		resultChan <- deletedResearch
	}()

	select {
	case err := <-errChan:
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao deletar pesquisa", "error": err})
		return
	case deletedResearch := <-resultChan:
		if len(deletedResearch) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "Nenhuma pesquisa encontrada para deletar"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Pesquisa apagada com sucesso."})
	}
}
