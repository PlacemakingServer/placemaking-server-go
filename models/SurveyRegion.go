package models

type SurveyRegion struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	LocationTitle string  `json:"location_title"`
	Lat           float32 `json:"lat"`
	Long          float32 `json:"long"`
	SurveyId      string  `json:"survey_id"`
	SurveyType    string  `json:"survey_type"`
}

type CreateSurveyRegion struct {
	Name          string  `json:"name"`
	Lat           float32 `json:"lat"`
	Long          float32 `json:"long"`
	LocationTitle string  `json:"location_title"`
}

type UpdateSurveyRegion struct {
	Name          string  `json:"name"`
	LocationTitle string  `json:"location_title"`
	Lat           float32 `json:"lat"`
	Long          float32 `json:"long"`
}
