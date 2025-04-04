package routes

import (
	"placemaking-backend-go/controllers"
	"placemaking-backend-go/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/auth/login", controllers.Login)
	router.POST("/auth/register", controllers.Register)

	api := router.Group("/api")
	api.Use(middleware.JWTAuthMiddleware())
	{
		//Rotas para Auth

		api.POST("/logout", controllers.Logout)


		//Rotas para Users
		api.GET("/users/:id", controllers.GetUserById)
		api.GET("/users", controllers.GetAllUsers)
		api.PUT("/users/:id", controllers.UpdateUserById)
		api.DELETE("/users/:id", controllers.DeleteUserById)

		//Rotas para Input_Types
		api.GET("/input_types", controllers.GetInputTypes)
	}

	return router
}
