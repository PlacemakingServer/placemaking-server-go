package repository

import (
	"encoding/json"
	"log"
	"placemaking-backend-go/config"
	"placemaking-backend-go/models"
)

func GetAllInputTypes() ([]models.InputType, error) {
	supabase := config.GetSupabase()

	// Executar a query para buscar todos os input_types
	response, _, err := supabase.From("input_types").Select("*", "", false).Execute()
	if err != nil {
		log.Println("Error fetching input types:", err)
		return nil, err
	}

	// Converter JSON para struct
	var inputTypes []models.InputType
	err = json.Unmarshal(response, &inputTypes) // <- Passamos `response` diretamente
	if err != nil {
		log.Println("Error decoding input types:", err)
		return nil, err
	}

	return inputTypes, nil
}
