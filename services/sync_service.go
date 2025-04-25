package services

import (
	"log"
	"time"

	"placemaking-backend-go/db"
	"placemaking-backend-go/models"
)

// Struct público para retorno seguro de usuários
type PublicUser struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func fetchSince[T any](table string, fromDate time.Time, filterByDate bool) ([]T, error) {
	var results []T
	query := db.GetSupabase().
		From(table).
		Select("*", "exact", false)

	if filterByDate {
		query = query.Gte("created_at", fromDate.Format(time.RFC3339))
	}

	_, err := query.ExecuteTo(&results)
	return results, err
}

func UpsertGeneric(table string, items []map[string]interface{}) error {
	_, _, err := db.GetSupabase().
		From(table).
		Upsert(items, "id", "minimal", "").
		Execute()

	return err
}

func GetUsersSince() ([]models.SanitizedUser, error) {
	supabase := db.GetSupabase()

	var users []models.User

	_, err := supabase.From("users").
		Select("*", "", false).
		ExecuteTo(&users)

	if err != nil {
		log.Println("Erro ao buscar usuários:", err)
		return nil, err
	}

	var sanitizedUsers []models.SanitizedUser
	for _, user := range users {
		user.ConvertTimestamps()
		sanitizedUsers = append(sanitizedUsers, models.SanitizeUser(user))
	}

	return sanitizedUsers, nil
}

func GetInputTypesSince(fromDate time.Time) ([]models.InputType, error) {
	return fetchSince[models.InputType]("input_types", fromDate, false)
}

func GetResearchesSince(fromDate time.Time) ([]models.Research, error) {
	return fetchSince[models.Research]("researches", fromDate, true)
}

func GetResearchContributorsSince(fromDate time.Time) ([]models.Contributor, error) {
	return fetchSince[models.Contributor]("research_contributors", fromDate, true)
}

func GetFieldsSince(fromDate time.Time) ([]models.Field, error) {
	return fetchSince[models.Field]("fields", fromDate, true)
}

func GetFieldOptionsSince(fromDate time.Time) ([]models.FieldOption, error) {
	return fetchSince[models.FieldOption]("field_options", fromDate, true)
}

func GetSurveyAnswersSince(fromDate time.Time) ([]models.SurveyAnswer, error) {
	return fetchSince[models.SurveyAnswer]("survey_answers", fromDate, true)
}

func GetStaticSurveysSince(fromDate time.Time) ([]models.Survey, error) {
	return fetchSince[models.Survey]("static_surveys", fromDate, true)
}

func GetFormSurveysSince(fromDate time.Time) ([]models.Survey, error) {
	return fetchSince[models.Survey]("form_surveys", fromDate, true)
}

func GetDynamicSurveysSince(fromDate time.Time) ([]models.Survey, error) {
	return fetchSince[models.Survey]("dynamic_surveys", fromDate, true)
}

func GetSurveyTimeRangesSince(fromDate time.Time) ([]models.SurveyTimeRange, error) {
	return fetchSince[models.SurveyTimeRange]("survey_time_ranges", fromDate, true)
}

func GetSurveyRegionsSince(fromDate time.Time) ([]models.SurveyRegion, error) {
	return fetchSince[models.SurveyRegion]("survey_regions", fromDate, true)
}

func GetSurveyGroupsSince(fromDate time.Time) ([]models.SurveyGroup, error) {
	return fetchSince[models.SurveyGroup]("survey_group", fromDate, true)
}

func GetSurveyContributorsSince(fromDate time.Time) ([]models.SurveyContributors, error) {
	return fetchSince[models.SurveyContributors]("survey_contributors", fromDate, true)
}
