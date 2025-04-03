package services

import (
	"placemaking-backend-go/models"
	repository "placemaking-backend-go/repositories"
)

func FetchAllActivityTypes () ([]models.ActivityType, error){
	return repository.GetAllActivityTypes()
}