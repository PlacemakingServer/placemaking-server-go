package routes

import (
	"github.com/gin-gonic/gin"
	"placemaking-backend-go/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		//Rotas para Users
		api.GET("/users/:id", controllers.GetUserById)

		//Rotas para Activity_Types
		api.GET("/activity_types", controllers.GetActivityTypes)

		//Rotas para Input_Types
		api.GET("/input_types", controllers.GetInputTypes)
	}

	return router
}
