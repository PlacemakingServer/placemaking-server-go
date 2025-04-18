package models

type SurveyTimeRange struct {
	ID         string `json:"id"`
	SurveyId   string `json:"survey_id"`
	SurveyType string `json:"survey_type"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
}

type CreateSurveyTimeRange struct {
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
}

type UpdateSurveyTimeRange struct {
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
}
