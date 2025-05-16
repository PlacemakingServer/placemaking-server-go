package repository

import (
	"log"
	"placemaking-backend-go/db"
	"placemaking-backend-go/models"
)
	// Função auxiliar para converter "" em nil
func ifEmptyNil(s string) interface{} {
	if s == "" {
		return nil
	}
	return s
}

func CreateSurveyAnswer(surveyId, surveyType, contributorId string, surveyAnswerData models.CreateSurveyAnswer) (models.SurveyAnswer, error) {
	supabase := db.GetSupabase()

	// Convertendo para map[string]interface{}
	insertData := map[string]interface{}{
		"value":                surveyAnswerData.Value,
		"survey_id":            surveyId,
		"survey_type":          surveyType,
		"survey_contributor_id":       contributorId,
		"registered_at":        surveyAnswerData.RegisteredAt,
		"survey_group_id":      surveyAnswerData.SurveyGroupId,
		"survey_time_range_id": ifEmptyNil(surveyAnswerData.SurveyTimeRangeId),
		"survey_region_id":     ifEmptyNil(surveyAnswerData.SurveyRegionId),
	}

	var createdSurveyAnswer models.SurveyAnswer

	_, err := supabase.
		From("survey_answers").
		Insert(insertData, false, "", "", "").
		Single().
		ExecuteTo(&createdSurveyAnswer)

	if err != nil {
		log.Println("[CreateSurveyAnswer] Erro ao criar resposta de pesquisa no banco de dados:", err)
		return models.SurveyAnswer{}, err
	}

	return createdSurveyAnswer, nil

}

func GetAllSurveyAnswersBySurveyId(surveyId string) ([]models.SurveyAnswer, error) {
	supabase := db.GetSupabase()

	var surveyAnswers []models.SurveyAnswer

	_, err := supabase.From("survey_answers").
		Select("*", "", false).
		Eq("survey_id", surveyId).
		ExecuteTo(&surveyAnswers)

	if err != nil {
		log.Println("[GetAllSurveyAnswersBySurveyId] Erro ao buscar respostas de pesquisa no banco de dados:", err)
		return []models.SurveyAnswer{}, err
	}

	return surveyAnswers, nil

}

func GetAllAnswersByContributorId(contributorId string) ([]models.SurveyAnswer, error) {
	supabase := db.GetSupabase()

	var surveyAnswers []models.SurveyAnswer

	_, err := supabase.From("survey_answers").
		Select("*", "", false).
		Eq("survey_contributor_id", contributorId).
		ExecuteTo(&surveyAnswers)

	if err != nil {
		log.Println("[GetAllAnswersByContributorId] Erro ao buscar respostas de pesquisa no banco de dados:", err)
		return []models.SurveyAnswer{}, err
	}

	return surveyAnswers, nil

}

func GetSurveyAnswerById(id string) (models.SurveyAnswer, error) {
	supabase := db.GetSupabase()

	var surveyAnswer models.SurveyAnswer

	_, err := supabase.From("survey_answers").
		Select("*", "", false).
		Eq("id", id).
		Single().
		ExecuteTo(&surveyAnswer)

	if err != nil {
		log.Println("[GetSurveyAnswerById] Erro ao buscar resposta de pesquisa no banco de dados:", err)
		return models.SurveyAnswer{}, err
	}

	return surveyAnswer, nil

}

func DeleteSurveyAnswerById(id string) ([]models.SurveyAnswer, error) {
	supabase := db.GetSupabase()

	var deletedAnswers []models.SurveyAnswer

	_, err := supabase.From("survey_answers").
		Delete("","").
		Eq("id", id).
		ExecuteTo(&deletedAnswers)

	if err != nil {
		log.Println("[DeleteSurveyAnswerById] Erro ao deletar resposta de pesquisa no banco de dados:", err)
		return deletedAnswers, err
	}

	return deletedAnswers, nil
}

func UpdateSurveyAnswerById(id string, surveyAnswerData models.UpdateSurveyAnswer) (models.SurveyAnswer, error) {
	supabase := db.GetSupabase()

	// Convertendo para map[string]interface{}
	updateData := map[string]interface{}{
		"value":      surveyAnswerData.Value,
		"registered_at": surveyAnswerData.RegisteredAt,
		"survey_group_id": surveyAnswerData.SurveyGroupId,
		"survey_time_range_id": surveyAnswerData.SurveyTimeRangeId,
		"survey_region_id": surveyAnswerData.SurveyRegionId,
	}

	var updatedSurveyAnswer models.SurveyAnswer

	_, err := supabase.
		From("survey_answers").
		Update(updateData, "", "").
		Eq("id", id).
		Single().
		ExecuteTo(&updatedSurveyAnswer)

	if err != nil {
		log.Println("[UpdateSurveyAnswerById] Erro ao atualizar resposta de pesquisa no banco de dados:", err)
		return models.SurveyAnswer{}, err
	}

	return updatedSurveyAnswer, nil

}