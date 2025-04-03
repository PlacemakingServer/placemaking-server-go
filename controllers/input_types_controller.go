package controllers

import (
    "net/http"
    "placemaking-backend-go/services"

    "github.com/gin-gonic/gin"
)

func GetInputTypes(c *gin.Context) {
	inputTypes, err := services.FetchInputTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch input types"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lista de tipos de entrada encontrados.","input_types": inputTypes})
}