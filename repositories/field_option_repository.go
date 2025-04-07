package repository

import (
	"log"
	"errors"
	"placemaking-backend-go/db"
	"placemaking-backend-go/models"
)

func CreateFieldOption(fieldId string, createFieldOptionData models.CreateFieldOption) (models.FieldOption, error) {
	supabase := db.GetSupabase()

	// Criando um mapa com os dados para inserção
	newFieldOption := map[string]interface{}{
		"field_id":     fieldId,
		"option_text":  createFieldOptionData.OptionText,
		"option_value": createFieldOptionData.OptionValue,
	}

	var fieldOption models.FieldOption

	_, err := supabase.From("field_options").
		Insert(newFieldOption, false, "", "", "").
		Single().
		ExecuteTo(&fieldOption)

	if err != nil {
		log.Println("[CreateFieldOption] Erro ao criar opção de campo:", err)
		return models.FieldOption{}, err
	}

	return fieldOption, nil
}

func GetAllFieldOptionsByFieldId(fieldId string) ([]models.FieldOption, error) {
	supabase := db.GetSupabase()

	var fieldOptions []models.FieldOption

	_, err := supabase.From("field_options").
		Select("*", "", false).
		Eq("field_id", fieldId).
		ExecuteTo(&fieldOptions)

	if err != nil {
		log.Println("[GetAllFieldOptionsByFieldId] Erro ao buscar opções de campo:", err)
		return nil, err
	}

	return fieldOptions, nil
}

func DeleteFieldOptionById(id, fieldId string) error {
	supabase := db.GetSupabase()

	var deletedField []models.FieldOption

	_, err := supabase.From("field_options").
		Delete("", "").
		Eq("id", id).
		Eq("field_id", fieldId).
		ExecuteTo(&deletedField)

	if err != nil {
		log.Println("[DeleteFieldOptionById] Erro ao deletar opção de campo:", err)
		return err
	}

	if len(deletedField) == 0 { // Verifica se o Supabase retornou algo
		log.Println("[DeleteFieldOptionById] Nenhuma opção de campo encontrada para deletar")
		return errors.New("nenhum campo encontrado para os critérios fornecidos")
	}

	return nil
}