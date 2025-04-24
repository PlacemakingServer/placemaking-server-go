package services

import (
	"time"

	"placemaking-backend-go/db"
	"placemaking-backend-go/models"
)

func fetchSince[T any](table string, fromDate time.Time) ([]T, error) {
	var results []T

	_, err := db.GetSupabase().
		From(table).
		Select("*", "exact", false).
		Gte("created_at", fromDate.Format(time.RFC3339)).
		ExecuteTo(&results)

	return results, err
}
func UpsertGeneric(table string, items []map[string]interface{}) error {
	_, _, err := db.GetSupabase().
		From(table).
		Upsert(items, "id", "minimal", "").
		Execute()

	return err
}

func GetUsersSince(fromDate time.Time) ([]models.User, error) {
	return fetchSince[models.User]("users", fromDate)
}

func GetResearchesSince(fromDate time.Time) ([]models.Research, error) {
	return fetchSince[models.Research]("researches", fromDate)
}

func GetResearchContributorsSince(fromDate time.Time) ([]models.Contributor, error) {
	return fetchSince[models.Contributor]("research_contributors", fromDate)
}

func GetFieldsSince(fromDate time.Time) ([]models.Field, error) {
	return fetchSince[models.Field]("fields", fromDate)
}

func GetInputTypesSince(fromDate time.Time) ([]models.InputType, error) {
	return fetchSince[models.InputType]("input_types", fromDate)
}

func GetFieldOptionsSince(fromDate time.Time) ([]models.FieldOption, error) {
	return fetchSince[models.FieldOption]("field_options", fromDate)
}

func GetSurveyAnswersSince(fromDate time.Time) ([]models.SurveyAnswer, error) {
	return fetchSince[models.SurveyAnswer]("survey_answers", fromDate)
}

func GetStaticSurveysSince(fromDate time.Time) ([]models.Survey, error) {
	return fetchSince[models.Survey]("static_surveys", fromDate)
}

func GetFormSurveysSince(fromDate time.Time) ([]models.Survey, error) {
	return fetchSince[models.Survey]("form_surveys", fromDate)
}

func GetDynamicSurveysSince(fromDate time.Time) ([]models.Survey, error) {
	return fetchSince[models.Survey]("dynamic_surveys", fromDate)
}

func GetSurveyTimeRangesSince(fromDate time.Time) ([]models.SurveyTimeRange, error) {
	return fetchSince[models.SurveyTimeRange]("survey_time_ranges", fromDate)
}

func GetSurveyRegionsSince(fromDate time.Time) ([]models.SurveyRegion, error) {
	return fetchSince[models.SurveyRegion]("survey_regions", fromDate)
}

func GetSurveyGroupsSince(fromDate time.Time) ([]models.SurveyGroup, error) {
	return fetchSince[models.SurveyGroup]("survey_group", fromDate)
}

func GetSurveyContributorsSince(fromDate time.Time) ([]models.SurveyContributors, error) {
	return fetchSince[models.SurveyContributors]("survey_contributors", fromDate)
}
