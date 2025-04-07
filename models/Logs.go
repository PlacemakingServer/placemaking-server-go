package models

type Logs struct {
	ID          string      `json:"id"`
	UserId      string      `json:"user_id"`
	Action      string      `json:"action"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
	CreatedAt   string      `json:"created_at"`
}

type CreateLogs struct {
	UserId      string      `json:"user_id"`
	Action      string      `json:"action"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}
