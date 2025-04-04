package models

type CreateResearch struct {
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	ReleaseDate   string  `json:"release_date"`
	CreatedBy     string  `json:"created_by"`
	Lat           float32 `json:"lat"`
	Long          float32 `json:"long"`
	LocationTitle string  `json:"location_title"`
	EndDate       string  `json:"end_date"`
}

type Research struct {
	Id            string  `json:"id"`
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	ReleaseDate   string  `json:"release_date"`
	CreatedBy     string  `json:"created_by"`
	Lat           float32 `json:"lat"`
	Long          float32 `json:"long"`
	LocationTitle string  `json:"location_title"`
	EndDate       string  `json:"end_date"`
}

type UpdateResearch struct {
	Title         string  `json:"title"`
	Description   string  `json:"description"`
	ReleaseDate   string  `json:"release_date"`
	Lat           float32 `json:"lat"`
	Long          float32 `json:"long"`
	LocationTitle string  `json:"location_title"`
	EndDate       string  `json:"end_date"`
}

