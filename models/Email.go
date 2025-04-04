package models

type EmailData struct {
	Name         string `json:"name"`
	TempPassword string `json:"temp_password"`
}

type RecoveryEmailData struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}
