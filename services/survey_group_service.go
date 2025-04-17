package services

import (
	"errors"
	"placemaking-backend-go/models"
	repository "placemaking-backend-go/repositories"
)

// GetSurveyGroups retorna todos os grupos de uma pesquisa específica
func GetSurveyGroups(surveyId, surveyType string) ([]models.SurveyGroup, error) {
	groups, err := repository.GetAllGroupsBySurveyId(surveyId, surveyType)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

// CreateSurveyGroup cria um novo grupo de pesquisa
func CreateSurveyGroup(surveyId string, data models.CreateSurveyGroup) (models.SurveyGroup, error) {
	// Validação básica dos campos
	if surveyId == "" || data.SurveyType == "" {
		return models.SurveyGroup{}, errors.New("survey_id ou survey_type não podem ser vazios")
	}

	group, err := repository.CreateSurveyGroup(surveyId, data)
	if err != nil {
		return models.SurveyGroup{}, err
	}
	return group, nil
}

// DeleteSurveyGroup remove um grupo de pesquisa pelo seu ID
func DeleteSurveyGroup(id string) error {
	if id == "" {
		return errors.New("id do grupo não pode ser vazio")
	}

	err := repository.DeleteSurveyGroup(id)
	if err != nil {
		return err
	}
	return nil
}
