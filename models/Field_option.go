package models

type FieldOption struct {
	ID          string `json:"id"`
	FieldId     string `json:"field_id"`
	OptionText  string `json:"option_text"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	OptionValue string `json:"option_value"`
}

type CreateFieldOption struct {
	OptionText  string `json:"option_text"`
	OptionValue string `json:"option_value"`
}