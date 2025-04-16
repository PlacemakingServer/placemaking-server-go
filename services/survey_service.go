package services

import (
	"errors"
	"placemaking-backend-go/models"
	repository "placemaking-backend-go/repositories"
)

// Criar uma nova pesquisa
func CreateSurvey(surveyType string, createSurveyData models.CreateSurvey) (models.Survey, error) {
	if createSurveyData.Title == "" || createSurveyData.ResearchId == "" {
		return models.Survey{}, errors.New("título e ID da pesquisa são obrigatórios")
	}

	return repository.CreateSurvey(surveyType, createSurveyData)
}

// Obter todas as pesquisas
func GetAllSurveys(surveyType string) ([]models.Survey, error) {
	if surveyType == "" {
		return nil, errors.New("o tipo de pesquisa é obrigatório")
	}

	return repository.GetAllSurveys(surveyType)
}

// Obter uma pesquisa por ID
func GetSurveyById(id, researchId, surveyType string) (models.Survey, error) {
	if id == "" {
		return models.Survey{}, errors.New("o ID da pesquisa é obrigatório")
	}

	return repository.GetSurveyById(id, researchId, surveyType)
}

// Atualizar uma pesquisa por ID
func UpdateSurveyById(id, surveyType string, updateData models.UpdateSurvey) (models.Survey, error) {
	if id == "" {
		return models.Survey{}, errors.New("o ID da pesquisa é obrigatório")
	}

	return repository.UpdateSurveyById(id, surveyType, updateData)
}

// Deletar uma pesquisa por ID
func DeleteSurveyById(id, surveyType string) (models.Survey, error) {
	if id == "" {
		return models.Survey{}, errors.New("o ID da pesquisa é obrigatório")
	}

	return repository.DeleteSurveyById(id, surveyType)
}

// Serviço para buscar pesquisas por research_id
func GetSurveysByResearchId(researchId, surveyType string) ([]models.Survey, error) {
	return repository.GetSurveysByResearchId(researchId, surveyType)
}
