package services

import (
    "errors"
    "placemaking-backend-go/models"
    repo "placemaking-backend-go/repositories"
)

// GetSurveyAnswers retorna todas as respostas de uma pesquisa específica
func GetSurveyAnswers(surveyId string) ([]models.SurveyAnswer, error) {
    if surveyId == "" {
        return nil, errors.New("surveyId não pode ser vazio")
    }

    answers, err := repo.GetAllAnswersBySurveyId(surveyId)
    if err != nil {
        return nil, err
    }
    return answers, nil
}

// CreateSurveyAnswer cria uma nova resposta para uma pesquisa
func CreateSurveyAnswer(surveyId, surveyType, contributorId string, data models.CreateSurveyAnswer) (models.SurveyAnswer, error) {
    // Validação básica dos campos
    if surveyId == "" || surveyType == "" || contributorId == "" {
        return models.SurveyAnswer{}, errors.New("surveyId, surveyType e contributorId não podem ser vazios")
    }
    if data.Value == "" {
        return models.SurveyAnswer{}, errors.New("o campo Value não pode ser vazio")
    }
    if data.SurveyGroupId == "" {
        return models.SurveyAnswer{}, errors.New("surveyGroupId não pode ser vazio")
    }

    answer, err := repo.CreateSurveyAnswer(surveyId, surveyType, contributorId, data)
    if err != nil {
        return models.SurveyAnswer{}, err
    }
    return answer, nil
}

// DeleteSurveyAnswer remove uma resposta de pesquisa pelo seu ID
func DeleteSurveyAnswer(id string) error {
    if id == "" {
        return errors.New("id da resposta não pode ser vazio")
    }

    err := repo.DeleteSurveyAnswer(id)
    if err != nil {
        return err
    }
    return nil
}
