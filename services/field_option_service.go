package services

import (
	"placemaking-backend-go/models"
	"placemaking-backend-go/repositories"
)

// CreateFieldOptionService cria uma nova opção para um campo
func CreateFieldOptionService(fieldId string, createFieldOptionData models.CreateFieldOption) (models.FieldOption, error) {
	return repository.CreateFieldOption(fieldId, createFieldOptionData)
}

// GetAllFieldOptionsByFieldIdService retorna todas as opções de um campo
func GetAllFieldOptionsByFieldIdService(fieldId string) ([]models.FieldOption, error) {
	return repository.GetAllFieldOptionsByFieldId(fieldId)
}

// DeleteFieldOptionByIdService deleta uma opção de campo por ID
func DeleteFieldOptionByIdService(id, fieldId string) error {
	return repository.DeleteFieldOptionById(id, fieldId)
}
