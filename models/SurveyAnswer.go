package models

type SurveyAnswer struct {
	ID            string `json:"id"`
	Value         string `json:"value"`
	SurveyType    string `json:"survey_type"`
	SurveyId      string `json:"survey_id"`
	SurveyGroupId string `json:"survey_group_id"`
	ContributorId string `json:"contributor_id"`
}

type CreateSurveyAnswer struct {
	Value         string `json:"value"`
	SurveyGroupId string `json:"survey_group_id"`
}