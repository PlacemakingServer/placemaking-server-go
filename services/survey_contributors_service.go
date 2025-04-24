package services

import (
	"placemaking-backend-go/models"
	"placemaking-backend-go/repositories"
	"log"
)

func CreateSurveyContributorService(id string, data models.CreateSurveyContributors) (models.SurveyContributors, error) {
	return repository.CreateSurveyContributors(id, data)
}

func GetSurveyContributorsBySurveyIdService(surveyId string) ([]models.ViewSurveyContributors, error) {
	contributors, err := repository.GetSurveyContributorsBySurveyId(surveyId)
	if err != nil {
		log.Println("[Service] Erro ao buscar survey contributors:", err)
		return nil, err
	}

	var viewContributors []models.ViewSurveyContributors
	for _, contributor := range contributors {
		user, err := repository.GetUserById(contributor.UserId)
		if err != nil {
			log.Println("[Service] Erro ao buscar usuário do survey contributor:", err)
			return nil, err
		}

		viewContributor := models.ViewSurveyContributors{
			ID:          contributor.ID,
			SurveyId:    contributor.SurveyId,
			SurveyType:  contributor.SurveyType,
			User:        models.SanitizeUser(user),
			Instruction: contributor.Instruction,
			CreatedAt:   contributor.CreatedAt,
			UpdatedAt:   contributor.UpdatedAt,
		}

		viewContributors = append(viewContributors, viewContributor)
	}

	return viewContributors, nil
}


func DeleteSurveyContributorsByIdService(id, surveyId string) error {
	return repository.DeleteSurveyContributorsById(id, surveyId)
}
