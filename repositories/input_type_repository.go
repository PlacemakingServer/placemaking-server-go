package repository

import (
	"encoding/json"
	"log"
	"placemaking-backend-go/db"
	"placemaking-backend-go/models"
)
func GetAllInputTypes() ([]models.InputType, error) {
	supabase := db.GetSupabase()

	// Executar a query para buscar todos os input_types
	response, _, err := supabase.From("input_types").Select("*", "", false).Execute()
	if err != nil {
		log.Println("Error fetching input types:", err)
		return nil, err
	}

	// Converter JSON para struct
	var inputTypes []models.InputType

	if err = json.Unmarshal(response, &inputTypes); err != nil {
		log.Println("Error decoding input types:", err)
		return nil, err
	}

	return inputTypes, nil
}

func GetInputTypeByID(id string) (*models.InputType, error) {
	supabase := db.GetSupabase()

	// Executar a query para buscar o input_type pelo ID
	response, _, err := supabase.From("input_types").Select("*", "", false).Eq("id", id).Single().Execute()
	if err != nil {
		log.Println("Error fetching input type by ID:", err)
		return nil, err
	}

	// Converter JSON para struct
	var inputType models.InputType

	if err = json.Unmarshal(response, &inputType); err != nil {
		log.Println("Error decoding input type by ID:", err)
		return nil, err
	}

	return &inputType, nil
}
