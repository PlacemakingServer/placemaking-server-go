package routes

import (
	"github.com/gin-gonic/gin"
	"placemaking-backend-go/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/activity_types", controllers.GetActivityTypes)
		api.GET("/input_types", controllers.GetInputTypes)
	}

	return router
}
