package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

// Settings armazena as configurações do .env
type Settings struct {
	Env         string
	ProjectName string
	SupabaseURL string
	SupabaseKey string
	GmailUser   string
	GmailPass   string
	JwtSecret   string
}

// LoadSettings carrega as variáveis do .env para a struct Settings
func LoadSettings() *Settings {
	err := godotenv.Load()
	if err != nil {
		log.Println("Erro ao carregar .env, usando variáveis do sistema.")
	}

	return &Settings{
		Env:         os.Getenv("ENV"),
		ProjectName: os.Getenv("PROJECT_NAME"),
		SupabaseURL: os.Getenv("SUPABASE_URL"),
		SupabaseKey: os.Getenv("SUPABASE_KEY"),
		GmailUser:   os.Getenv("GMAIL_USER"),
		GmailPass:   os.Getenv("GMAIL_PASSWORD"),
		JwtSecret:   os.Getenv("JWT_SECRET"),
	}
}