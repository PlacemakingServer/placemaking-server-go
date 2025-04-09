package services

import (
	"log"
	"placemaking-backend-go/models"
	repository "placemaking-backend-go/repositories"
)

func CreateContributor(researchId, userId string, contributorData models.CreateContributor) (models.Contributor, error) {
	contributor, err := repository.CreateContributor(researchId, userId, contributorData)
	if err != nil {
		log.Println("[Service] Erro ao criar colaborador:", err)
		return models.Contributor{}, err
	}
	return contributor, nil
}

func GetContributorById(id string) (models.Contributor, error) {
	contributor, err := repository.GetContributorById(id)
	if err != nil {
		log.Println("[Service] Erro ao buscar colaborador por ID:", err)
		return models.Contributor{}, err
	}
	return contributor, nil
}

func UpdateContributorById(id string, updateData models.UpdateContributor) (models.Contributor, error) {
	contributor, err := repository.UpdateContributorById(id, updateData)
	if err != nil {
		log.Println("[Service] Erro ao atualizar colaborador:", err)
		return models.Contributor{}, err
	}
	return contributor, nil
}

func DeleteContributorById(id string) error {
	err := repository.DeleteContributorById(id)
	if err != nil {
		log.Println("[Service] Erro ao deletar colaborador:", err)
		return err
	}
	return nil
}

func GetAllContributorsByResearchId(researchId string) ([]models.ViewContributor, error) {
	contributors, err := repository.GetAllContributorsByResearchId(researchId)
	if err != nil {
		log.Println("[Service] Erro ao buscar colaboradores por ID da pesquisa:", err)
		return nil, err
	}

	var viewContributors []models.ViewContributor
    for _, contributor := range contributors {
        user, err := repository.GetUserById(contributor.UserId) // Supondo que essa função exista
        if err != nil {
            log.Println("[Service] Erro ao buscar usuário do colaborador:", err)
            return nil, err
        }

        viewContributor := models.ViewContributor{
            ID:          contributor.ID,
            ResearchId:  contributor.ResearchId,
            UserId:      contributor.UserId,
            Instruction: contributor.Instruction,
            User:        models.SanitizeUser(user), // Atribui o usuário à estrutura
        }
        viewContributors = append(viewContributors, viewContributor)
    }

	return viewContributors, nil
}

func GetContributorByResearchAndUserId(researchId, userId string) (models.Contributor, error) {
	contributor, err := repository.GetContributorByResearchAndUserId(researchId, userId)
	if err != nil {
		log.Println("[Service] Erro ao buscar colaborador por pesquisa e usuário:", err)
		return models.Contributor{}, err
	}
	return contributor, nil
}

func DeleteContributorByResearchAndUserId(researchId, userId string) error {
	err := repository.DeleteContributorByResearchAndUserId(researchId, userId)
	if err != nil {
		log.Println("[Service] Erro ao deletar colaborador por pesquisa e usuário:", err)
		return err
	}
	return nil
}
