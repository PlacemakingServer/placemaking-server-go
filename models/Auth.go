package models

type AuthEmailRecovery struct {
	Email string `json:"email"`
}

type AuthValidadetToken struct {
	Token string `json:"token"`
}

type AuthUserResetPassword struct {
	NewPassword string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}
