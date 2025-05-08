package models

type Contributor struct {
	ID          string `json:"id"`
	ResearchId  string `json:"research_id"`
	UserId      string `json:"user_id"`
	Instruction string `json:"instruction"`
}

type ContributorSanitize struct {
	ID          string `json:"id"`
	UserId      string `json:"user_id"`
	Instruction string `json:"instruction"`
}

type ViewContributor struct {
	ID          string        `json:"id"`
	ResearchId  string        `json:"research_id"`
	UserId      string        `json:"user_id"`
	Instruction string        `json:"instruction"`
	User        SanitizedUser `json:"user"`
}

type CreateContributor struct {
	Instruction string `json:"instruction"`
}

type UpdateContributor struct {
	Instruction string `json:"instruction"`
}
