package routes

import (
	"placemaking-backend-go/controllers"
	"placemaking-backend-go/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Lista de rotas públicas dentro de /api/v1
	publicRoutes := map[string]bool{
		"POST /api/v1/auth/login":           true,
		"POST /api/v1/auth/register":        true,
		"POST /api/v1/auth/forgot_password": true,
		"POST /api/v1/auth/validate_code":   true,
	}

	// Todas as rotas dentro de /api/v1 passam pelo middleware
	api := router.Group("/api/v1")
	api.Use(middleware.JWTAuthMiddleware(publicRoutes))
	{
		// Rotas de autenticação
		api.POST("/auth/login", controllers.Login)
		api.POST("/auth/register", controllers.Register)
		api.POST("/auth/forgot_password", controllers.ForgotPassword)
		api.POST("/auth/validate_code", controllers.ValidateCode)

		// Rotas protegidas
		api.POST("/auth/logout", controllers.Logout)
		api.PUT("/auth/reset_password", controllers.ResetPassword)

		// Rotas de usuários
		api.GET("/users/:id", controllers.GetUserById)
		api.GET("/users", controllers.GetAllUsers)
		api.PUT("/users/:id", controllers.UpdateUserById)
		api.DELETE("/users/:id", controllers.DeleteUserById)

		// Rotas de tipos de input
		api.GET("/input_types", controllers.GetInputTypes)
	}

	return router
}
