package repository

import (
	"errors"
	"log"
	"placemaking-backend-go/db"
	"placemaking-backend-go/models"
)

func CreateSurveyContributors(surveyId string, createSurveyContributorData models.CreateSurveyContributors) (models.SurveyContributors, error) {
	supabase := db.GetSupabase()

	// Criando o mapa com os dados para inserção
	new_data := map[string]interface{}{
		"survey_id":   surveyId,
		"survey_type": createSurveyContributorData.SurveyType,
		"user_id":     createSurveyContributorData.UserId,
		"instruction": createSurveyContributorData.Instruction,
	}

	var surveyContributor models.SurveyContributors

	// Inserindo no banco de dados
	_, err := supabase.From("survey_contributors").
		Insert(new_data, false, "", "", "").
		Single().
		ExecuteTo(&surveyContributor)

	if err != nil {
		return models.SurveyContributors{}, err
	}

	return surveyContributor, nil
}

func GetSurveyContributorsBySurveyId(surveyId string) ([]models.SurveyContributors, error) {
	supabase := db.GetSupabase()

	var surveyContributors []models.SurveyContributors

	_, err := supabase.From("survey_contributors").
		Select("*", "", false).
		Eq("survey_id", surveyId).
		ExecuteTo(&surveyContributors)

	if err != nil {
		log.Println("[GetSurveyContributorsBySurveyId] Erro ao buscar contribuidores.", err)
		return []models.SurveyContributors{}, err
	}

	return surveyContributors, nil
}

func DeleteSurveyContributorsById(user_id, surveyId string) error {
	supabase := db.GetSupabase()

	var deletedContributor []models.SurveyContributors

	_, err := supabase.From("survey_contributors").
		Delete("", "").
		Eq("user_id", user_id).
		Eq("survey_id", surveyId).
		ExecuteTo(&deletedContributor)

	if err != nil {
		log.Println("[DeleteSurveyContributorsById]: Erro ao apagar contribuidor da pesquisa.", err)
	}

	if len(deletedContributor) == 0 {
		log.Println("[DeleteSurveyContributorsById]: Nenhum colaborador encontrado.")
		return errors.New("erro ao buscar jogador no banco de dados")
	}

	return nil
}
