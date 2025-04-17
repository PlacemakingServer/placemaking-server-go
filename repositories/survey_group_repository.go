package repository

import (
	"errors"
	"log"
	"placemaking-backend-go/db"
	"placemaking-backend-go/models"
)

func GetAllGroupsBySurveyId(surveyId, surveyType string) ([]models.SurveyGroup, error) {

	supabase := db.GetSupabase()

	var surveyGroups []models.SurveyGroup

	_, err := supabase.From("survey_group").Select("*","", false).Eq("survey_id", surveyId).Eq("survey_type", surveyType).ExecuteTo(&surveyGroups)

	if err != nil {
		log.Println("[GetAllGroupsBySurveyId] Erro ao buscar grupo no banco de dados")
		return []models.SurveyGroup{}, err
	}

	return surveyGroups, nil

}

func CreateSurveyGroup(surveyId string, surveyData models.CreateSurveyGroup) (models.SurveyGroup, error) {
	supabase := db.GetSupabase()

	// Convertendo para map[string]interface{}
	insertData := map[string]interface{}{
		"survey_id":   surveyId,
		"survey_type": surveyData.SurveyType,
	}

	var createdGroup models.SurveyGroup

	_, err := supabase.
		From("survey_group").
		Insert(insertData,false,"","","").
		Single().
		ExecuteTo(&createdGroup)

	if err != nil {
		log.Println("[CreateSurveyGroup] Erro ao criar grupo no banco de dados:", err)
		return models.SurveyGroup{}, err
	}

	return createdGroup, nil
}

func DeleteSurveyGroup(id string) error {
    supabase := db.GetSupabase()

    // Executa o delete na tabela survey_group, filtrando pelo ID

	var res []models.SurveyGroup

    _, err := supabase.
        From("survey_group").
        Delete("","").
        Eq("id", id).
        ExecuteTo(res)

    if err != nil {
        log.Printf("[DeleteSurveyGroup] Erro ao deletar grupo: %v\n", err)
        return err
    }

	if len(res) == 0 {
		return errors.New("survey Id selecionado não existe ou já deletado")
	}
	
	return nil
}