package models

type SurveyGroup struct {
	ID         string `json:"id"`
	SurveyId   string `json:"survey_id"`
	SurveyType string `json:"survey_type"`
}

type CreateSurveyGroup struct {
	SurveyId   string `json:"survey_id"`
	SurveyType string `json:"survey_type"`
}
