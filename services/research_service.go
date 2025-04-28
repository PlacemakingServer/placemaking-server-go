package services

import (
	"log"
	"placemaking-backend-go/models"
	repository "placemaking-backend-go/repositories"
)

func FetchCreateResearch(createResearchData models.CreateResearch) (models.Research, error) {
	research, err := repository.CreateResearch(createResearchData)

	if err != nil {
		log.Println("[FetchCreateResearch] Erro ao criar pesquisa:", err)
		return models.Research{}, err
	}

	return research, nil
}

func FetchAllResearches() ([]models.Research, error) {

	researches, err := repository.GetAllResearches()
	if err != nil {
		log.Println("[Service] Erro ao buscar pesquisas:", err)
		return nil, err
	}

	// var viewResearches []models.ViewResearch
	// for _, research := range researches {
	// 	user, err := repository.GetUserById(research.CreatedBy)
	// 	if err != nil {
	// 		log.Println("[Service] Erro ao buscar usuário da pesquisa:", err)
	// 		return nil, err
	// 	}

	// 	viewResearch := models.ViewResearch{
	// 		Id:            research.Id,
	// 		Title:         research.Title,
	// 		Description:   research.Description,
	// 		ReleaseDate:   research.ReleaseDate,
	// 		CreatedBy:     models.SanitizeUser(user),
	// 		Lat:           research.Lat,
	// 		Long:          research.Long,
	// 		LocationTitle: research.LocationTitle,
	// 		EndDate:       research.EndDate,
	// 	}

	// 	viewResearches = append(viewResearches, viewResearch)
	// }

	return researches, nil
}

func FetchResearchById(id string) (models.ViewResearch, error) {

	research, err := repository.GetResearchById(id)
	if err != nil {
		log.Println("[Service] Erro ao buscar pesquisa por ID:", err)
		return models.ViewResearch{}, err
	}

	user, err := repository.GetUserById(research.CreatedBy)
	if err != nil {
		log.Println("[Service] Erro ao buscar usuário da pesquisa:", err)
		return models.ViewResearch{}, err
	}

	viewResearch := models.ViewResearch{
		Id:            research.Id,
		Title:         research.Title,
		Description:   research.Description,
		ReleaseDate:   research.ReleaseDate,
		CreatedBy:     models.SanitizeUser(user),
		Lat:           research.Lat,
		Long:          research.Long,
		LocationTitle: research.LocationTitle,
		EndDate:       research.EndDate,
	}

	return viewResearch, nil
}

func FetchUpdateResearch(id string, updateResearchData models.UpdateResearch) (models.Research, error) {
	research, err := repository.UpdateResearchById(id, updateResearchData)
	if err != nil {
		log.Println("[FetchUpdateResearch] Erro ao atualizar pesquisa:", err)
		return models.Research{}, err
	}
	return research, nil
}

func FetchDeleteResearch(id string) ([]models.Research, error) {
	deletedResearch, err := repository.DeleteResearchById(id)
	if err != nil {
		log.Println("[FetchDeleteResearch] Erro ao deletar pesquisa:", err)
		return deletedResearch, err
	}
	return deletedResearch, nil
}
