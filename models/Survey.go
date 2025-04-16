package models

type Survey struct {
	ID            string  `json:"id"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	Lat           float32 `json:"lat"`
	Long          float32 `json:"long"`
	LocationTitle string  `json:"location_title"`
	ResearchId    string  `json:"research_id"`
}

type CreateSurvey struct {
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	Lat           float32 `json:"lat"`
	Long          float32 `json:"long"`
	LocationTitle string  `json:"location_title"`
	ResearchId    string  `json:"research_id"`
}

type UpdateSurvey struct {
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	Lat           float32 `json:"lat"`
	Long          float32 `json:"long"`
	LocationTitle string  `json:"location_title"`
}

type SurveyType struct {
	Type string `json:"survey_type"`
}