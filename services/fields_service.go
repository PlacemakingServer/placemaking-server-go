package services

import (
	"placemaking-backend-go/models"
	"placemaking-backend-go/repositories"
)

// CreateFieldService cria um novo campo para um survey específico
func CreateFieldService(surveyId, surveyType string, createFieldData models.CreateField) (models.Field, error) {
	return repository.CreateField(surveyId, surveyType, createFieldData)
}

// GetAllFieldsBySurveyIdService retorna todos os campos de um survey específico
func GetAllFieldsBySurveyIdService(surveyId, surveyType string) ([]models.Field, error) {
	return repository.GetAllFieldsBySurveyId(surveyId, surveyType)
}

// DeleteFieldBySurveyIdService deleta um campo específico de um survey
func DeleteFieldBySurveyIdService(id, surveyId, surveyType string) error {
	return repository.DeleteFieldBySurveyId(id, surveyId, surveyType)
}

// UpdateFieldService atualiza um campo específico de um survey
func UpdateFieldService(id, surveyId, surveyType string, updateFieldData models.CreateField) (models.Field, error) {
	return repository.UpdateField(id, surveyId, surveyType, updateFieldData)
}
