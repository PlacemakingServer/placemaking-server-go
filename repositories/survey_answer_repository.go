package repository

import (
	"errors"
	"log"
	"placemaking-backend-go/db"
	"placemaking-backend-go/models"
)

func GetAllAnswersBySurveyId(surveyId string) ([]models.SurveyAnswer, error) {
	supabase := db.GetSupabase()

	var surveyAnswers []models.SurveyAnswer

	_, err := supabase.From("survey_answers").Select("*", "", false).Eq("survey_id", surveyId).ExecuteTo(&surveyAnswers)

	if err != nil {
		log.Println("[GetAllAnswersBySurveyId] Erro ao buscar respostas no banco de dados:", err)
		return []models.SurveyAnswer{}, err
	}

	return surveyAnswers, nil
}

func CreateSurveyAnswer(surveyId, surveyType, contributorId string, answerData models.CreateSurveyAnswer) (models.SurveyAnswer, error) {
	supabase := db.GetSupabase()

	// Convertendo para map[string]interface{}
	insertData := map[string]interface{}{
		"value":      answerData.Value,
		"survey_group_id": answerData.SurveyGroupId,
		"survey_id":  surveyId,
		"survey_type": surveyType,
		"contributor_id": contributorId,
	}

	var createdAnswer models.SurveyAnswer

	_, err := supabase.
		From("survey_answers").
		Insert(insertData, false, "", "", "").
		Single().
		ExecuteTo(&createdAnswer)

	if err != nil {
		log.Println("[CreateSurveyAnswer] Erro ao criar resposta no banco de dados:", err)
		return models.SurveyAnswer{}, err
	}

	return createdAnswer, nil
}

func DeleteSurveyAnswer(id string) error {
	supabase := db.GetSupabase()

	// Executa o delete na tabela survey_answers, filtrando pelo ID
	res, _, err := supabase.
		From("survey_answers").
		Delete("", "").
		Eq("id", id).
		Execute()

	if err != nil {
		log.Printf("[DeleteSurveyAnswer] Erro ao deletar resposta: %v\n", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("resposta selecionada não existe ou já foi deletada")
	}

	return nil
}
