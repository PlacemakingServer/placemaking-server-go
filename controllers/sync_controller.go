package controllers

import (
	"net/http"
	"time"

	"placemaking-backend-go/services"

	"github.com/gin-gonic/gin"
)

func SyncGet(c *gin.Context) {
	entity := c.Param("entity")
	if entity == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entidade não especificada"})
		return
	}

	fromDate := time.Now().AddDate(0, 0, -20)

	var data interface{}
	var err error

	switch entity {
	case "users":
		data, err = services.GetUsersSince(fromDate)
	case "researches":
		data, err = services.GetResearchesSince(fromDate)
	case "research_contributors":
		data, err = services.GetResearchContributorsSince(fromDate)
	case "fields":
		data, err = services.GetFieldsSince(fromDate)
	case "input_types":
		data, err = services.GetInputTypesSince(fromDate)
	case "field_options":
		data, err = services.GetFieldOptionsSince(fromDate)
	case "survey_answers":
		data, err = services.GetSurveyAnswersSince(fromDate)
	case "static_surveys":
		data, err = services.GetStaticSurveysSince(fromDate)
	case "form_surveys":
		data, err = services.GetFormSurveysSince(fromDate)
	case "dynamic_surveys":
		data, err = services.GetDynamicSurveysSince(fromDate)
	case "survey_time_ranges":
		data, err = services.GetSurveyTimeRangesSince(fromDate)
	case "survey_regions":
		data, err = services.GetSurveyRegionsSince(fromDate)
	case "survey_group":
		data, err = services.GetSurveyGroupsSince(fromDate)
	case "survey_contributors":
		data, err = services.GetSurveyContributorsSince(fromDate)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entidade não reconhecida"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{entity: data})
}
