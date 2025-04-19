package repository

import (
	"log"
	"placemaking-backend-go/db"
	"placemaking-backend-go/models"
)

func CreateSurveyTimeRange(surveyId, surveyType string, surveyTimeData models.CreateSurveyTimeRange) (models.SurveyTimeRange, error) {
	supabase := db.GetSupabase()

	// Convertendo para map[string]interface{}
	insertData := map[string]interface{}{
		"start_time": surveyTimeData.StartTime,
		"end_time":   surveyTimeData.EndTime,
		"survey_id":  surveyId,
		"survey_type": surveyType,
	}

	var createdSurveyTimeRange models.SurveyTimeRange

	_, err := supabase.
		From("survey_time_ranges").
		Insert(insertData, false, "", "", "").
		Single().
		ExecuteTo(&createdSurveyTimeRange)

	if err != nil {
		log.Println("[CreateSurveyTimeRange] Erro ao criar intervalo de tempo de pesquisa no banco de dados:", err)
		return models.SurveyTimeRange{}, err
	}

	return createdSurveyTimeRange, nil
}

func GetAllSurveyRegionsBySurveyType(surveyType string) ([]models.SurveyTimeRange, error) {
	supabase := db.GetSupabase()

	var surveyTimeRanges []models.SurveyTimeRange

	_, err := supabase.From("survey_time_ranges").
		Select("*", "", false).
		Eq("survey_type", surveyType).
		ExecuteTo(&surveyTimeRanges)

	if err != nil {
		log.Println("[GetAllSurveyRegionsBySurveyType] Erro ao buscar regiões de pesquisa no banco de dados:", err)
		return []models.SurveyTimeRange{}, err
	}

	return surveyTimeRanges, nil
}

func GetAllSurveyTimeRangeBySurveyId(surveyId string) ([]models.SurveyTimeRange, error) {
	supabase := db.GetSupabase()

	var surveyTimeRanges []models.SurveyTimeRange

	_, err := supabase.From("survey_time_ranges").
		Select("*", "", false).
		Eq("survey_id", surveyId).
		ExecuteTo(&surveyTimeRanges)

	if err != nil {
		log.Println("[GetSurveyTimeRangeBySurveyId] Erro ao buscar intervalos de tempo de pesquisa no banco de dados:", err)
		return []models.SurveyTimeRange{}, err
	}

	return surveyTimeRanges, nil
}

func GetSurveyTimeRangeById(id string) (models.SurveyTimeRange, error) {
	supabase := db.GetSupabase()

	var surveyTimeRange models.SurveyTimeRange

	_, err := supabase.From("survey_time_ranges").
		Select("*", "", false).
		Eq("id", id).
		Single().
		ExecuteTo(&surveyTimeRange)

	if err != nil {
		log.Println("[GetSurveyTimeRangeById] Erro ao buscar intervalo de tempo de pesquisa no banco de dados:", err)
		return models.SurveyTimeRange{}, err
	}

	return surveyTimeRange, nil
}

func UpdateSurveyTimeRange(id string, surveyTimeData models.UpdateSurveyTimeRange) (models.SurveyTimeRange, error) {
	supabase := db.GetSupabase()

	// Convertendo para map[string]interface{}
	updateData := map[string]interface{}{
		"start_time": surveyTimeData.StartTime,
		"end_time":   surveyTimeData.EndTime,
	}

	var updatedSurveyTimeRange models.SurveyTimeRange

	_, err := supabase.
		From("survey_time_ranges").
		Update(updateData, "", "").
		Eq("id", id).
		Single().
		ExecuteTo(&updatedSurveyTimeRange)

	if err != nil {
		log.Println("[UpdateSurveyTimeRange] Erro ao atualizar intervalo de tempo de pesquisa no banco de dados:", err)
		return models.SurveyTimeRange{}, err
	}

	return updatedSurveyTimeRange, nil
}

func DeleteSurveyTimeRange(id string) ([]models.SurveyTimeRange, error) {
	supabase := db.GetSupabase()

	var deletedSurveyTimeRange []models.SurveyTimeRange

	// Executa o delete na tabela survey_time_ranges, filtrando pelo ID
	_, err := supabase.
		From("survey_time_ranges").
		Delete("", "").
		Eq("id", id).
		ExecuteTo(&deletedSurveyTimeRange)

	if err != nil {
		log.Println("[DeleteSurveyTimeRange] Erro ao deletar intervalo de tempo de pesquisa no banco de dados:", err)
		return deletedSurveyTimeRange, err
	}

	return deletedSurveyTimeRange, nil
}