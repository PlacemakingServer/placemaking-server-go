package services

import (
	"log"
	"placemaking-backend-go/models"
	repository "placemaking-backend-go/repositories"
)

func CreateSurveyTimeRange(surveyId, surveyType string, data models.CreateSurveyTimeRange) (models.SurveyTimeRange, error) {
	surveyTimeRange, err := repository.CreateSurveyTimeRange(surveyId, surveyType, data)
	if err != nil {
		log.Println("[Service] Erro ao criar intervalo de tempo de pesquisa:", err)
		return models.SurveyTimeRange{}, err
	}
	return surveyTimeRange, nil
}

func GetAllSurveyTimeRangeBySurveyId(surveyId string) ([]models.SurveyTimeRange, error) {
	surveyTimeRanges, err := repository.GetAllSurveyTimeRangeBySurveyId(surveyId)
	if err != nil {
		log.Println("[Service] Erro ao buscar intervalos de tempo por surveyId:", err)
		return nil, err
	}
	return surveyTimeRanges, nil
}

func GetAllSurveyTimeRangeBySurveyType(surveyType string) ([]models.SurveyTimeRange, error) {
	surveyTimeRanges, err := repository.GetAllSurveyRegionsBySurveyType(surveyType)
	if err != nil {
		log.Println("[Service] Erro ao buscar intervalos de tempo por surveyType:", err)
		return nil, err
	}
	return surveyTimeRanges, nil
}

func GetSurveyTimeRangeById(id string) (models.SurveyTimeRange, error) {
	surveyTimeRange, err := repository.GetSurveyTimeRangeById(id)
	if err != nil {
		log.Println("[Service] Erro ao buscar intervalo de tempo por ID:", err)
		return models.SurveyTimeRange{}, err
	}
	return surveyTimeRange, nil
}

func UpdateSurveyTimeRange(id string, data models.UpdateSurveyTimeRange) (models.SurveyTimeRange, error) {
	updatedSurveyTimeRange, err := repository.UpdateSurveyTimeRange(id, data)
	if err != nil {
		log.Println("[Service] Erro ao atualizar intervalo de tempo:", err)
		return models.SurveyTimeRange{}, err
	}
	return updatedSurveyTimeRange, nil
}

func DeleteSurveyTimeRange(id string) ([]models.SurveyTimeRange, error) {
	deleted, err := repository.DeleteSurveyTimeRange(id)
	if err != nil {
		log.Println("[Service] Erro ao deletar intervalo de tempo:", err)
		return nil, err
	}
	return deleted, nil
}
