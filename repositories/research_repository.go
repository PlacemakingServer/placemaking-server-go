package repository

import (
	"log"
	"placemaking-backend-go/db"
	"placemaking-backend-go/models"
	"time"
)

func CreateResearch(createResearchData models.CreateResearch) (models.Research, error) {

	supabase := db.GetSupabase()

	var research models.Research

	// Inserção no banco de dados
	_, err := supabase.From("researches").Insert(map[string]interface{}{
		"title":          createResearchData.Title,
		"description":    createResearchData.Description,
		"release_date":   createResearchData.ReleaseDate,
		"created_by":     createResearchData.CreatedBy,
		"lat":            createResearchData.Lat,
		"long":           createResearchData.Long,
		"location_title": createResearchData.LocationTitle,
		"end_date":       createResearchData.EndDate,
	}, false, "", "", "").Single().ExecuteTo(&research)

	if err != nil {
		return models.Research{}, err
	}

	return research, nil
}

func GetAllResearches() ([]models.ContributorsResearch, error) {
	supabase := db.GetSupabase()

	var researches []models.ContributorsResearch

	_, err := supabase.From("researches").
		Select("*, research_contributors(id, user_id, instruction)", "", false). // join com a tabela de contribuidores
		Eq("status", "true").
		ExecuteTo(&researches)

	if err != nil {
		log.Println("Erro ao buscar pesquisas com contribuidores:", err)
		return nil, err
	}

	return researches, nil
}


func GetResearchById(id string) (models.Research, error) {
	supabase := db.GetSupabase()

	var research models.Research

	_, err := supabase.From("researches").
		Select("*", "", false).
		Single().
		Eq("id", id).
		ExecuteTo(&research)

	if err != nil {
		log.Println("Erro ao buscar pesquisas:", err)
		return models.Research{}, err
	}

	return research, nil

}

func UpdateResearchById(id string, updateResearchData models.UpdateResearch) (models.Research, error) {
	supabase := db.GetSupabase()

	updatedData := map[string]interface{}{
		"title":          updateResearchData.Title,
		"description":    updateResearchData.Description,
		"release_date":   updateResearchData.ReleaseDate,
		"end_date":       updateResearchData.EndDate,
		"lat":            updateResearchData.Lat,
		"long":           updateResearchData.Long,
		"location_title": updateResearchData.LocationTitle,
		"updated_at":     time.Now(), // Atualiza a data de modificação
		"status":         updateResearchData.Status,
	}

	var research models.Research

	_, err := supabase.From("researches").
		Update(updatedData, "", ""). // Atualiza os dados no Supabase
		Eq("id", id).                // Filtra pelo ID
		Single().
		ExecuteTo(&research) // Decodifica para a struct Research

	if err != nil {
		log.Println("Erro ao atualizar pesquisa:", err)
		return models.Research{}, err
	}

	return research, nil
}

func DeleteResearchById(id string) ([]models.Research, error) {
	supabase := db.GetSupabase()

	var deletedResearch []models.Research

	_, err := supabase.From("researches").
		Delete("", "").
		Eq("id", id).
		ExecuteTo(&deletedResearch)

	if err != nil {
		log.Println("Erro ao deletar pesquisa:", err)
		return deletedResearch, err
	}

	return deletedResearch, nil
}
