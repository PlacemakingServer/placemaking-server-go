package repository

import (
	"log"
	"placemaking-backend-go/db"
	"placemaking-backend-go/models"
)

func CreateSurvey(surveyType string, createSurveyData models.CreateSurvey) (models.Survey, error) {
	supabase := db.GetSupabase()

	// Criando um mapa para representar os dados da pesquisa
	newSurvey := map[string]interface{}{
		"title":          createSurveyData.Title,
		"description":    createSurveyData.Description,
		"lat":            createSurveyData.Lat,
		"long":           createSurveyData.Long,
		"location_title": createSurveyData.LocationTitle,
		"research_id":    createSurveyData.ResearchId,
	}

	var survey models.Survey

	// Inserindo no banco de dados
	_, err := supabase.From(surveyType).
		Insert(newSurvey, false, "", "", "").
		ExecuteTo(&survey)

	if err != nil {
		log.Println("[CreateSurvey] Erro ao criar pesquisa:", err)
		return models.Survey{}, err
	}

	return survey, nil
}

func GetAllSurveys(surveyType string) ([]models.Survey, error) {
	supabase := db.GetSupabase()

	var surveys []models.Survey

	_, err := supabase.From(surveyType).
		Select("*", "", false).
		ExecuteTo(&surveys)

	if err != nil {
		log.Println("[GetAllSurveys] Erro ao buscar pesquisas:", err)
		return nil, err
	}

	return surveys, nil
}

func GetSurveyById(id, surveyType string) (models.Survey, error) {
	supabase := db.GetSupabase()

	var survey models.Survey

	_, err := supabase.From(surveyType).
		Select("*", "", false).
		Eq("id", id).
		Single().
		ExecuteTo(&survey)

	if err != nil {
		log.Println("[GetSurveyById] Erro ao buscar pesquisa:", err)
		return models.Survey{}, err
	}

	return survey, nil
}

func UpdateSurveyById(id, surveyType string, updateData models.UpdateResearch) (models.Survey, error) {
	supabase := db.GetSupabase()

	// Criando um mapa com os dados atualizados
	updatedFields := map[string]interface{}{
		"title":          updateData.Title,
		"description":    updateData.Description,
		"lat":            updateData.Lat,
		"long":           updateData.Long,
		"location_title": updateData.LocationTitle,
	}

	var survey models.Survey

	_, err := supabase.From(surveyType).
		Update(updatedFields, "", "").
		Eq("id", id).
		Single().
		ExecuteTo(&survey)

	if err != nil {
		log.Println("[UpdateSurveyById] Erro ao atualizar pesquisa:", err)
		return models.Survey{}, err
	}

	return survey, nil
}

func DeleteSurveyById(id, surveyType string) (models.Survey, error) {
	supabase := db.GetSupabase()

	var survey models.Survey

	_, _, err := supabase.From(surveyType).
		Delete("","").
		Eq("id", id).
		Single().
		Execute()

	if err != nil {
		log.Println("[DeleteSurveyById] Erro ao deleatar pesquisa:", err)
		return models.Survey{}, err
	}

	return survey, nil
}

// Buscar pesquisas por research_id
func GetSurveysByResearchId(researchId, surveyType string) ([]models.Survey, error) {
	supabase := db.GetSupabase()

	var surveys []models.Survey

	_, err := supabase.From(surveyType).
		Select("*", "", false).
		Eq("research_id", researchId).
		ExecuteTo(&surveys)

	if err != nil {
		return nil, err
	}

	return surveys, nil
}
