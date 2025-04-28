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
				contributors.POST("/:userId", controllers.CreateContributor)
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
			contributor.GET("/:contributorId", controllers.GetContributorById)
			contributor.PUT("/:contributorId", controllers.UpdateContributorById)
			contributor.DELETE("/:contributorId", controllers.DeleteContributorById)
			contributor.GET("/:contributorId/answers", controllers.GetSurveyAnswersByContributorId)
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

			survey_contributors := survey.Group("/:surveyId/contributors")
			{
				survey_contributors.GET("", controllers.GetSurveyContributorsBySurveyId)
				survey_contributors.POST("", controllers.CreateSurveyContributor)
				survey_contributors.DELETE("/:contributorId", controllers.DeleteSurveyContributor)
			}

			survey_answers := survey.Group("/:surveyId/answers")
			{
				survey_answers.POST("", controllers.CreateSurveyAnswer)
				survey_answers.GET("", controllers.GetSurveyAnswersBySurveyId)
				survey_answers.DELETE("/:answerId", controllers.DeleteSurveyAnswerById)
				survey_answers.GET("/:answerId", controllers.GetSurveyAnswerById)
				survey_answers.PUT("/:answerId", controllers.UpdateSurveyAnswerById)
			}

			survey_group := survey.Group("/:surveyId/group")
			{
				survey_group.POST("", controllers.CreateSurveyGroup)
				survey_group.GET("", controllers.GetSurveyGroups)
				survey_group.DELETE("/:groupId", controllers.DeleteSurveyGroup)
			}

			survey_region := survey.Group("/:surveyId/region")
			{
				survey_region.POST("", controllers.CreateSurveyRegion)
				survey_region.GET("", controllers.GetAllSurveyRegionsBySurveyId)
				survey_region.DELETE("/:regionId", controllers.DeleteSurveyRegion)
				survey_region.GET("/:regionId", controllers.GetSurveyRegionById)
				survey_region.PUT("/:regionId", controllers.UpdateSurveyRegionById)

			}

			survey_time_range := survey.Group("/:surveyId/time")
			{
				survey_time_range.POST("", controllers.CreateSurveyTimeRange)
				survey_time_range.GET("", controllers.GetAllSurveyTimeRangeBySurveyId)
				survey_time_range.DELETE("/:timeRangeId", controllers.DeleteSurveyTimeRange)
				survey_time_range.GET("/:timeRangeId", controllers.GetSurveyTimeRangeById)
				survey_time_range.PUT("/:timeRangeId", controllers.UpdateSurveyTimeRangeById)
			}
		}
		//Grupo de Field options
		field_option := api.Group("/fields/:fieldId/options")
		{
			field_option.GET("", controllers.GetAllFieldOptionsByFieldId)
			field_option.POST("", controllers.CreateFieldOption)
			field_option.DELETE("/:optionId", controllers.DeleteFieldOptionById)
		}

		// Grupo de sincronização
		// sync := api.Group("/sync")
		// {
		// 	sync.GET("/:entity", controllers.SyncGet)
		// 	sync.PATCH("/:entity", controllers.SyncPatch)
		// }

		// Grupo de tipos de input
		api.GET("/input_types", controllers.GetInputTypes)
	}

	return router
}
