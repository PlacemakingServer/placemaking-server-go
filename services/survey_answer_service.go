package services

import (
	"log"
	"placemaking-backend-go/models"
	repository "placemaking-backend-go/repositories"
)

// CreateSurveyAnswer cria uma nova resposta de pesquisa no banco de dados
func CreateSurveyAnswer(surveyId, surveyType, contributorId string, data models.CreateSurveyAnswer) (models.SurveyAnswer, error) {

	answer, err := repository.CreateSurveyAnswer(surveyId, surveyType, contributorId, data)
	if err != nil {
		log.Println("[CreateSurveyAnswer] Erro ao criar resposta de pesquisa no banco de dados:", err)
		return models.SurveyAnswer{}, err
	}
	return answer, nil
}

func GetSurveyAnswersBySurveyId(surveyId string) ([]models.SurveyAnswer, error) {
	answers, err := repository.GetAllSurveyAnswersBySurveyId(surveyId)
	if err != nil {
		log.Println("[GetSurveyAnswersBySurveyId] Erro ao buscar respostas de pesquisa no banco de dados:", err)
		return nil, err
	}
	return answers, nil
}

func GetSurveyAnswersByContributorId(contributorId string) ([]models.SurveyAnswer, error) {
	answers, err := repository.GetAllAnswersByContributorId(contributorId)
	if err != nil {
		log.Println("[GetSurveyAnswersByContributorId] Erro ao buscar respostas de pesquisa no banco de dados:", err)
		return nil, err
	}
	return answers, nil
}

func GetSurveyAnswerById(id string) (models.SurveyAnswer, error) {
	answer, err := repository.GetSurveyAnswerById(id)
	if err != nil {
		log.Println("[GetSurveyAnswerById] Erro ao buscar resposta de pesquisa no banco de dados:", err)
		return models.SurveyAnswer{}, err
	}
	return answer, nil
}

func UpdateSurveyAnswerById(id string, data models.UpdateSurveyAnswer) (models.SurveyAnswer, error) {
	answers, err := repository.UpdateSurveyAnswerById(id, data)
	if err != nil {
		log.Println("[UpdateSurveyAnswerById] Erro ao atualizar resposta de pesquisa no banco de dados:", err)
		return models.SurveyAnswer{}, err
	}
	return answers, nil
}

func DeleteSurveyAnswerById(id string) ([]models.SurveyAnswer, error) {
	answers, err := repository.DeleteSurveyAnswerById(id)
	if err != nil {
		log.Println("[DeleteSurveyAnswerById] Erro ao deletar resposta de pesquisa no banco de dados:", err)
		return nil, err
	}
	return answers, nil
}