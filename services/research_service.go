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
	var researchs []models.Research

	researchs, err := repository.GetAllResearches()

	if err != nil {
		log.Println("[FetchAllResearchs] Erro ao buscar pesquisas:", err)
	}

	return researchs, nil
}

func FetchResearchById(id string) (models.Research, error){
	
	research, err := repository.GetResearchById(id)

	if err != nil {
		log.Println("[FecthResearchById] Erro ao buscar pesquisa:", err)
	}

	return research, nil
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
