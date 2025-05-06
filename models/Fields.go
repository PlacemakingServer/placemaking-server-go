package models

type Field struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	InputTypeId string `json:"input_type_id"`
	SurveyType  string `json:"survey_type"`
	SurveyId    string `json:"survey_id"`
	Position    int    `json:"position"`
}

type CreateField struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	InputTypeId string `json:"input_type_id"`
	Position    int    `json:"position"`
}
