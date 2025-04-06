package models

type Field struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Descriptiom string `json:"description"`
	InputTypeId string `json:"input_type_id"`
	SurveyType  string `json:"survey_type"`
	SurveyId    string `json:"survey_id"`
}

type CreateField struct {
	Title       string `json:"title"`
	Descriptiom string `json:"description"`
	InputTypeId string `json:"input_type_id"`
	SurveyType  string `json:"survey_type"`
	SurveyId    string `json:"survey_id"`
}
