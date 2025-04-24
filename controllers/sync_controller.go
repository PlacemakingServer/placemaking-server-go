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

func SyncPatch(c *gin.Context) {
	entity := c.Param("entity")
	if entity == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entidade não especificada"})
		return
	}

	var items []map[string]interface{}
	if err := c.BindJSON(&items); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	var err error

	switch entity {
	case "users":
		err = services.UpsertGeneric("users", items)
	case "researches":
		err = services.UpsertGeneric("researches", items)
	case "research_contributors":
		err = services.UpsertGeneric("research_contributors", items)
	case "fields":
		err = services.UpsertGeneric("fields", items)
	case "input_types":
		err = services.UpsertGeneric("input_types", items)
	case "field_options":
		err = services.UpsertGeneric("field_options", items)
	case "survey_answers":
		err = services.UpsertGeneric("survey_answers", items)
	case "static_surveys":
		err = services.UpsertGeneric("static_surveys", items)
	case "form_surveys":
		err = services.UpsertGeneric("form_surveys", items)
	case "dynamic_surveys":
		err = services.UpsertGeneric("dynamic_surveys", items)
	case "survey_time_ranges":
		err = services.UpsertGeneric("survey_time_ranges", items)
	case "survey_regions":
		err = services.UpsertGeneric("survey_regions", items)
	case "survey_group":
		err = services.UpsertGeneric("survey_group", items)
	case "survey_contributors":
		err = services.UpsertGeneric("survey_contributors", items)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Entidade não reconhecida"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
