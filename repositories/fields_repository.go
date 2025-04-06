package repository

import (
	"log"
	"placemaking-backend-go/db"
	"placemaking-backend-go/models"
)

func CreateField(surveyId, surveyType string, createFieldData models.CreateField) (models.Field, error) {
	supabase := db.GetSupabase()

	var field models.Field

	// Criando um novo field com os dados recebidos
	data := map[string]interface{}{
		"title":         createFieldData.Title,
		"description":   createFieldData.Description,
		"input_type_id": createFieldData.InputTypeId,
		"survey_id":     surveyId,
		"survey_type":   surveyType,
	}

	_, err := supabase.From("fields").
		Insert(data, false, "", "", "").
		Single().
		ExecuteTo(&field)

	if err != nil {
		log.Println("[CreateField] Erro ao criar campo:", err)
		return models.Field{}, err
	}

	return field, nil
}

func GetAllFieldsBySurveyId(surveyId, surveyType string) ([]models.Field, error) {
	supabase := db.GetSupabase()

	var fields []models.Field

	_, err := supabase.From("fields").
		Select("*", "", false).
		Eq("survey_id", surveyId).
		Eq("survey_type", surveyType).
		ExecuteTo(&fields)

	if err != nil {
		log.Println("[GetAllFieldsBySurveyId] Erro ao buscar campos:", err)
		return []models.Field{}, err
	}

	return fields, nil
}

func DeleteFieldBySurveyId(id, surveyId, surveyType string) error {
	supabase := db.GetSupabase()

	_, _, err := supabase.From("fields").
		Delete("", "").
		Eq("id", id).
		Eq("survey_id", surveyId).
		Eq("survey_type", surveyType).
		Execute()

	if err != nil {
		log.Println("[DeleteFieldBySurveyId] Erro ao deletar campo:", err)
		return err
	}

	return nil
}

func UpdateField(id, surveyId, surveyType string, updateFieldData models.CreateField) (models.Field, error) {
	supabase := db.GetSupabase()

	var field models.Field

	// Atualizando field com os dados recebidos
	data := map[string]interface{}{
		"title":         updateFieldData.Title,
		"description":   updateFieldData.Description,
		"input_type_id": updateFieldData.InputTypeId,
	}

	_, err := supabase.From("fields").
		Update(data, "", "").
		Eq("id", id).
		Eq("survey_id", surveyId).
		Eq("survey_type", surveyType).
		Single().
		ExecuteTo(&field)

	if err != nil {
		log.Println("[Update] Erro ao atualizar campo:", err)
		return models.Field{}, err
	}

	return field, nil
}