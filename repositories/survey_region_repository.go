package repository

import (
	"log"
	"placemaking-backend-go/db"
	"placemaking-backend-go/models"
)

func CreateSurveyRegion(surveyId, surveyType string, surveyRegionData models.CreateSurveyRegion) (models.SurveyRegion, error) {
	supabase := db.GetSupabase()

	// Convertendo para map[string]interface{}
	insertData := map[string]interface{}{
		"name":        surveyRegionData.Name,
		"lat":         surveyRegionData.Lat,
		"location_title": surveyRegionData.LocationTitle,
		"long":        surveyRegionData.Long,
		"survey_id":   surveyId,
		"survey_type": surveyType,
	}

	var createdSurveyRegion models.SurveyRegion

	_, err := supabase.
		From("survey_regions").
		Insert(insertData, false, "", "", "").
		Single().
		ExecuteTo(&createdSurveyRegion)

	if err != nil {
		log.Println("[CreateSurveyRegion] Erro ao criar região de pesquisa no banco de dados:", err)
		return models.SurveyRegion{}, err
	}

	return createdSurveyRegion, nil

}

func GetAllSurveyRegionsBySurveyId(surveyId string) ([]models.SurveyRegion, error) {
	supabase := db.GetSupabase()

	var surveyRegions []models.SurveyRegion

	_, err := supabase.From("survey_regions").
		Select("*", "", false).
		Eq("survey_id", surveyId).
		ExecuteTo(&surveyRegions)

	if err != nil {
		log.Println("[GetAllSurveyRegionsBySurveyId] Erro ao buscar regiões de pesquisa no banco de dados:", err)
		return []models.SurveyRegion{}, err
	}

	return surveyRegions, nil

}

func GetAllSurveysBySurveyType(surveyType string) ([]models.SurveyRegion, error) {
	supabase := db.GetSupabase()

	var surveyRegions []models.SurveyRegion

	_, err := supabase.From("survey_regions").
		Select("*", "", false).
		Eq("survey_type", surveyType).
		ExecuteTo(&surveyRegions)

	if err != nil {
		log.Println("[GetAllSurveysBySurveyType] Erro ao buscar regiões de pesquisa no banco de dados:", err)
		return []models.SurveyRegion{}, err
	}

	return surveyRegions, nil

}

func GetSurveyRegionById(id string) (models.SurveyRegion, error) {
	supabase := db.GetSupabase()

	var surveyRegion models.SurveyRegion

	_, err := supabase.From("survey_regions").
		Select("*", "", false).
		Eq("id", id).
		Single().
		ExecuteTo(&surveyRegion)

	if err != nil {
		log.Println("[GetSurveyRegionById] Erro ao buscar região de pesquisa no banco de dados:", err)
		return models.SurveyRegion{}, err
	}

	return surveyRegion, nil

}

func UpdateSurveyRegion(id string, surveyRegionData models.UpdateSurveyRegion) (models.SurveyRegion, error) {
	supabase := db.GetSupabase()

	// Convertendo para map[string]interface{}
	updateData := map[string]interface{}{
		"name": surveyRegionData.Name,
		"lat":  surveyRegionData.Lat,
		"long": surveyRegionData.Long,
		"location_title": surveyRegionData.LocationTitle,
	}

	var updatedSurveyRegion models.SurveyRegion

	_, err := supabase.
		From("survey_regions").
		Update(updateData, "", "").
		Eq("id", id).
		Single().
		ExecuteTo(&updatedSurveyRegion)

	if err != nil {
		log.Println("[UpdateSurveyRegion] Erro ao atualizar região de pesquisa no banco de dados:", err)
		return models.SurveyRegion{}, err
	}

	return updatedSurveyRegion, nil

}

func DeleteSurveyRegion(id string) ([]models.SurveyRegion, error) {
	supabase := db.GetSupabase()

	var deletedRecords []models.SurveyRegion

	// Executa o delete na tabela survey_region, filtrando pelo ID
	_, err := supabase.
		From("survey_regions").
		Delete("", "").
		Eq("id", id).
		ExecuteTo(&deletedRecords)

	if err != nil {
		log.Println("[DeleteSurveyRegion] Erro ao deletar região de pesquisa:", err)
		return deletedRecords, err
	}

	return deletedRecords, nil
}
