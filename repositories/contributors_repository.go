package repository

import (
	"log"
	"fmt"
	"placemaking-backend-go/db"
	"placemaking-backend-go/models"
)

func CreateContributor(researchId string, contributorData models.CreateContributor) (models.Contributor, error) {
	supabase := db.GetSupabase()

	var contributor models.Contributor

	// Criando um mapa para representar o colaborador
	newContributor := map[string]interface{}{
		"user_id":     contributorData.UserId,
		"research_id": researchId,
		"instruction": contributorData.Instruction,
	}

	// Inserindo no banco de dados
	_, err := supabase.From("research_contributors").
	Insert(newContributor, false, "", "", "").
	Single().
	ExecuteTo(&contributor)

	if err != nil {
		log.Println("[CreateContributor] Erro ao criar colaborador:", err)
		return models.Contributor{}, err
	}

	return contributor, nil
}

func GetContributorById(id string) (models.Contributor, error) {
	supabase := db.GetSupabase()

	var contributor models.Contributor

	_, err := supabase.From("research_contributors").
	Select("*","", false).
	Eq("id", id).
	Single().
	ExecuteTo(&contributor)

	if err != nil {
		log.Println("[GetContributorById] Erro ao buscar colaborador:", err)
		return models.Contributor{}, err
	}

	return contributor, nil
}

func UpdateContributorById(id string, updateData models.UpdateContributor) (models.Contributor, error){
	supabase := db.GetSupabase()

	var contributor models.Contributor

	// Criando um mapa para representar o colaborador
	newContributor := map[string]interface{}{
		"instruction": updateData.Instruction,
	}

	// Atualizando no banco de dados
	_, err := supabase.From("research_contributors").
	Update(newContributor,"","").
	Single().
	ExecuteTo(&contributor)

	if err != nil {
		log.Println("[UpdateContributorById] Erro ao atualizar contribuidor:", err)
		return models.Contributor{}, err
	}

	return contributor, nil
}

func DeleteContributorById(id string) error {
	supabase := db.GetSupabase()

	var existingContributors []models.Contributor

	// Verificar se o colaborador existe no banco
	_, err := supabase.From("research_contributors").
		Select("*", "", false).
		Eq("id", id).
		ExecuteTo(&existingContributors)

	if err != nil {
		log.Println("[DeleteContributorById] Erro ao buscar contribuidor:", err)
		return err
	}

	// Se não existir, retorna um erro
	if len(existingContributors) == 0 {
		return fmt.Errorf("colaborador não encontrado")
	}

	// Se existir, deleta o colaborador
	_, _, err = supabase.From("research_contributors").
		Delete("", "").
		Eq("id", id).
		Execute()

	if err != nil {
		log.Println("[DeleteContributorById] Erro ao deletar contribuidor no banco:", err)
		return err
	}

	return nil
}


func GetAllContributorsByResearchId(researchId string) ([]models.Contributor, error) {
	supabase := db.GetSupabase()

	var contributors []models.Contributor

	_, err := supabase.From("research_contributors").
	Select("*", "", false).
	Eq("research_id", researchId).
	ExecuteTo(&contributors)

	if err != nil {
		log.Println("[GetAllContributorsByResearchId] Erro ao buscar contribuidores no banco:", err)
		return []models.Contributor{}, err
	}

	return contributors, nil
}

func GetContributorByResearchAndUserId(researchId, userId string) (models.Contributor, error) {
	supabase := db.GetSupabase()

	var contributor models.Contributor

	_, err := supabase.From("research_contributors").
	Select("*", "", false).
	Eq("research_id", researchId).
	Eq("user_id", userId).
	Single().
	ExecuteTo(&contributor)

	if err != nil {
		log.Println("[GetContributorByResearchAndUserId] Erro ao buscar contribuidor no banco:", err)
		return models.Contributor{}, err
	}

	return contributor, nil
}

func DeleteContributorByResearchAndUserId(researchId, userId string) error {
	supabase := db.GetSupabase()

	var existingContributors []models.Contributor

	// Verificar se o colaborador existe no banco
	_, err := supabase.From("research_contributors").
		Select("*", "", false).
		Eq("research_id", researchId).
		Eq("user_id", userId).
		ExecuteTo(&existingContributors)

	if err != nil {
		log.Println("[DeleteContributorByResearchAndUserId] Erro ao buscar contribuidor:", err)
		return err
	}

	// Se não existir, retorna um erro
	if len(existingContributors) == 0 {
		return fmt.Errorf("colaborador não encontrado")
	}

	// Se existir, deleta o colaborador
	_, _, err = supabase.From("research_contributors").
		Delete("", "").
		Eq("research_id", researchId).
		Eq("user_id", userId).
		Execute()

	if err != nil {
		log.Println("[DeleteContributorByResearchAndUserId] Erro ao deletar contribuidor no banco:", err)
		return err
	}

	return nil
}
