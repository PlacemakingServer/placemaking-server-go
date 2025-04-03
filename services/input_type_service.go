package services

import (
    "placemaking-backend-go/models"
    "placemaking-backend-go/repositories"
)

func FetchInputTypes() ([]models.InputType, error) {
    return repository.GetAllInputTypes()
}
