package repository

import (
	"log"
	"fmt"
	"placemaking-backend-go/db"
	"placemaking-backend-go/models"
)

var surveyTypeMap = map[string]string{
	"Formulário": "form_surveys",
	"Estática":   "static_surveys",
	"Dinâmica":   "dynamic_surveys",
}

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

	tableName, exists := surveyTypeMap[surveyType]
	if !exists {
		err := fmt.Errorf("tipo de survey inválido: %s", surveyType)
		log.Println("[CreateSurvey] Erro:", err)
		return models.Survey{}, err
	}
	
	// Inserindo no banco de dados
	_, err := supabase.From(tableName).
		Insert(newSurvey, false, "", "", "").
		Single().
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

	tableName, exists := surveyTypeMap[surveyType]
	if !exists {
		err := fmt.Errorf("tipo de survey inválido: %s", surveyType)
		log.Println("[CreateSurvey] Erro:", err)
		return []models.Survey{}, err
	}

	_, err := supabase.From(tableName).
		Select("*", "", false).
		ExecuteTo(&surveys)

	if err != nil {
		log.Println("[GetAllSurveys] Erro ao buscar pesquisas:", err)
		return nil, err
	}

	return surveys, nil
}

func GetSurveyById(id, researchId, surveyType string) (models.Survey, error) {
	supabase := db.GetSupabase()

	tableName, exists := surveyTypeMap[surveyType]
	if !exists {
		err := fmt.Errorf("tipo de survey inválido: %s", surveyType)
		log.Println("[CreateSurvey] Erro:", err)
		return models.Survey{}, err
	}
	

	var survey models.Survey

	_, err := supabase.From(tableName).
		Select("*", "", false).
		Eq("id", id).
		Eq("research_id", researchId).
		Single().
		ExecuteTo(&survey)

	if err != nil {
		log.Println("[GetSurveyById] Erro ao buscar pesquisa:", err)
		return models.Survey{}, err
	}

	return survey, nil
}

func UpdateSurveyById(id, surveyType string, updateData models.UpdateSurvey) (models.Survey, error) {
	supabase := db.GetSupabase()

	tableName, exists := surveyTypeMap[surveyType]
	if !exists {
		err := fmt.Errorf("tipo de survey inválido: %s", surveyType)
		log.Println("[CreateSurvey] Erro:", err)
		return models.Survey{}, err
	}

	// Criando um mapa com os dados atualizados
	updatedFields := map[string]interface{}{
		"title":          updateData.Title,
		"description":    updateData.Description,
		"lat":            updateData.Lat,
		"long":           updateData.Long,
		"location_title": updateData.LocationTitle,
	}

	var survey models.Survey

	_, err := supabase.From(tableName).
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

	tableName, exists := surveyTypeMap[surveyType]
	if !exists {
		err := fmt.Errorf("tipo de survey inválido: %s", surveyType)
		log.Println("[CreateSurvey] Erro:", err)
		return models.Survey{}, err
	}


	var survey models.Survey

	_, _, err := supabase.From(tableName).
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

	tableName, exists := surveyTypeMap[surveyType]
	if !exists {
		err := fmt.Errorf("tipo de survey inválido: %s", surveyType)
		log.Println("[CreateSurvey] Erro:", err)
		return []models.Survey{}, err
	}

	var surveys []models.Survey

	_, err := supabase.From(tableName).
		Select("*", "", false).
		Eq("research_id", researchId).
		ExecuteTo(&surveys)

	if err != nil {
		return nil, err
	}

	return surveys, nil
}
