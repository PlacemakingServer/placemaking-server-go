package services

import (
	"log"
	"placemaking-backend-go/models"
	repository "placemaking-backend-go/repositories"
)

// Criar uma nova região vinculada a uma pesquisa
func CreateSurveyRegion(surveyId, surveyType string, regionData models.CreateSurveyRegion) (models.SurveyRegion, error) {
	surveyRegion, err := repository.CreateSurveyRegion(surveyId, surveyType, regionData)
	if err != nil {
		log.Println("[Service] Erro ao criar região de pesquisa:", err)
		return models.SurveyRegion{}, err
	}
	return surveyRegion, nil
}

// Buscar todas as regiões vinculadas a uma pesquisa pelo ID
func GetAllSurveyRegionsBySurveyId(surveyId string) ([]models.SurveyRegion, error) {
	surveyRegions, err := repository.GetAllSurveyRegionsBySurveyId(surveyId)
	if err != nil {
		log.Println("[Service] Erro ao buscar regiões por surveyId:", err)
		return nil, err
	}
	return surveyRegions, nil
}

// Buscar todas as regiões de pesquisas de um tipo específico
func GetAllSurveysBySurveyType(surveyType string) ([]models.SurveyRegion, error) {
	surveyRegions, err := repository.GetAllSurveysBySurveyType(surveyType)
	if err != nil {
		log.Println("[Service] Erro ao buscar regiões por surveyType:", err)
		return nil, err
	}
	return surveyRegions, nil
}

// Buscar uma região específica pelo ID
func GetSurveyRegionById(id string) (models.SurveyRegion, error) {
	surveyRegion, err := repository.GetSurveyRegionById(id)
	if err != nil {
		log.Println("[Service] Erro ao buscar região por ID:", err)
		return models.SurveyRegion{}, err
	}
	return surveyRegion, nil
}

// Atualizar uma região de pesquisa
func UpdateSurveyRegion(id string, data models.UpdateSurveyRegion) (models.SurveyRegion, error) {
	updatedRegion, err := repository.UpdateSurveyRegion(id, data)
	if err != nil {
		log.Println("[Service] Erro ao atualizar região de pesquisa:", err)
		return models.SurveyRegion{}, err
	}
	return updatedRegion, nil
}

// Deletar uma região de pesquisa
func DeleteSurveyRegion(id string) ([]models.SurveyRegion, error) {
	deleted, err := repository.DeleteSurveyRegion(id)
	if err != nil {
		log.Println("[Service] Erro ao deletar região de pesquisa:", err)
		return nil, err
	}
	return deleted, nil
}
