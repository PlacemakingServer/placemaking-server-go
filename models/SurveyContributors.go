package models

type SurveyContributors struct {
	ID          string `json:"id"`
	SurveyId    string `json:"survey_id"`
	SurveyType  string `json:"survey_type"`
	UserId      string `json:"user_id"`
	Instruction string `json:"instruction"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type CreateSurveyContributors struct {
	SurveyType  string `json:"survey_type"`
	UserId      string `json:"user_id"`
	Instruction string `json:"instruction"`
}
