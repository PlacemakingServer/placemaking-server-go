package controllers

import (
    "net/http"
    "placemaking-backend-go/services"

    "github.com/gin-gonic/gin"
)

func GetActivityTypes(c *gin.Context) {
	activity_types, err := services.FetchAllActivityTypes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error to fetch all activity types"})
		return
	}
	c.JSON(http.StatusOK, gin.H {"message": "Lista de tipos coleta encontrados.", "activity_types": activity_types})
}	