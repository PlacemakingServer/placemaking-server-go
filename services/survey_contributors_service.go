package services

import (
	"placemaking-backend-go/models"
	"placemaking-backend-go/repositories"
)

func CreateSurveyContributorService(id string, data models.CreateSurveyContributors) (models.SurveyContributors, error) {
	return repository.CreateSurveyContributors(id, data)
}

func GetSurveyContributorsBySurveyIdService(surveyId string) ([]models.SurveyContributors, error) {
	return repository.GetSurveyContributorsBySurveyId(surveyId)
}

func DeleteSurveyContributorsByIdService(id, surveyId string) error {
	return repository.DeleteSurveyContributorsById(id, surveyId)
}
