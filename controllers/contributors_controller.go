package controllers

import (
	"log"
	"net/http"
	"placemaking-backend-go/models"
	"placemaking-backend-go/services"

	"github.com/gin-gonic/gin"
)

// Criar um novo colaborador
func CreateContributor(c *gin.Context) {
	var contributorData models.CreateContributor
	if err := c.ShouldBindJSON(&contributorData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Dados inválidos", "error": err.Error()})
		return
	}
	researchId := c.Param("researchId")

	contributor, err := services.CreateContributor(researchId, contributorData)
	if err != nil {
		log.Println("[CreateContributor] Erro ao criar colaborador:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao criar colaborador", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Colaborador criado com sucesso", "contributor": contributor})
}

// Buscar colaborador por ID
func GetContributorById(c *gin.Context) {
	id := c.Param("id")

	contributor, err := services.GetContributorById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Colaborador não encontrado", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Colaborador encontrado", "contributor": contributor})
}

// Atualizar colaborador por ID
func UpdateContributorById(c *gin.Context) {
	id := c.Param("id")

	var updateData models.UpdateContributor
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Dados inválidos", "error": err.Error()})
		return
	}

	contributor, err := services.UpdateContributorById(id, updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao atualizar colaborador", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Colaborador atualizado com sucesso", "contributor": contributor})
}

// Deletar colaborador por ID
func DeleteContributorById(c *gin.Context) {
	id := c.Param("id")

	err := services.DeleteContributorById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao deletar colaborador", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Colaborador deletado com sucesso"})
}

// Buscar todos os colaboradores de uma pesquisa
func GetAllContributorsByResearchId(c *gin.Context) {
	researchId := c.Param("researchId")

	contributors, err := services.GetAllContributorsByResearchId(researchId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao buscar colaboradores", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lista de colaboradores obtida com sucesso", "contributors": contributors})
}

// Buscar colaborador por pesquisa e usuário
func GetContributorByResearchAndUserId(c *gin.Context) {
	researchId := c.Param("researchId")
	userId := c.Param("userId")

	contributor, err := services.GetContributorByResearchAndUserId(researchId, userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Colaborador não encontrado", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Colaborador encontrado", "contributor": contributor})
}

// Deletar colaborador por pesquisa e usuário
func DeleteContributorByResearchAndUserId(c *gin.Context) {
	researchId := c.Param("researchId")
	userId := c.Param("userId")

	err := services.DeleteContributorByResearchAndUserId(researchId, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao deletar colaborador", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Colaborador removido com sucesso"})
}
