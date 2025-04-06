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
		// Grupo de autenticação
		auth := api.Group("/auth")
		{
			auth.POST("/login", controllers.Login)
			auth.POST("/register", controllers.Register)
			auth.POST("/forgot_password", controllers.ForgotPassword)
			auth.POST("/validate_code", controllers.ValidateCode)
			auth.POST("/logout", controllers.Logout)
			auth.PUT("/reset_password", controllers.ResetPassword)
		}

		// Grupo de usuários
		users := api.Group("/users")
		{
			users.GET("/:id", controllers.GetUserById)
			users.GET("", controllers.GetAllUsers)
			users.PUT("/:id", controllers.UpdateUserById)
			users.DELETE("/:id", controllers.DeleteUserById)
		}

		// Grupo de pesquisas
		researches := api.Group("/research")
		{
			researches.POST("", controllers.CreateResearch)
			researches.GET("", controllers.GetAllResearches)
			researches.GET("/:researchId", controllers.GetResearchById)
			researches.PUT("/:researchId", controllers.UpdateResearch)
			researches.DELETE("/:researchId", controllers.DeleteResearch)

			// Grupo de colaboradores dentro de uma pesquisa (evitando conflito com :id)
			contributors := researches.Group("/:researchId/contributors")
			{
				contributors.POST("", controllers.CreateContributor)
				contributors.GET("", controllers.GetAllContributorsByResearchId)
				contributors.GET("/:userId", controllers.GetContributorByResearchAndUserId)
				contributors.DELETE("/:userId", controllers.DeleteContributorByResearchAndUserId)
			}

			// Grupo de surveys dentro de uma pesquisa
			surveys := researches.Group("/:researchId/survey")
			{
				surveys.GET("", controllers.GetSurveysByResearchId)
				surveys.GET("/:surveyId", controllers.GetSurveyById) //survey**
				surveys.DELETE("/:surveyId", controllers.DeleteSurveyById)
			}
		}

		//Rotas para Contributors
		contributor := api.Group("/contributors")
		{
			contributor.GET("/:id", controllers.GetContributorById)
			contributor.PUT("/:id", controllers.UpdateContributorById)
			contributor.DELETE("/:id", controllers.DeleteContributorById)
		}

		//Rotas para surveys
		survey := api.Group("/survey")
		{
			survey.POST("", controllers.CreateSurvey)
			survey.PUT("/:surveyId", controllers.UpdateSurveyById)

			fields := survey.Group("/:surveyId/fields")
				{
					fields.POST("", controllers.CreateField)
					fields.GET("", controllers.GetAllFieldsBySurveyId)
					fields.PUT("/:fieldId", controllers.UpdateField)
					fields.DELETE("/:fieldId", controllers.DeleteField)
				}

		}

		// Grupo de tipos de input
		api.GET("/input_types", controllers.GetInputTypes)
	}

	return router
}
