package services

import (
    "placemaking-backend-go/models"
    repository "placemaking-backend-go/repositories"
)

func FetchInputTypes() ([]models.InputType, error) {
    return repository.GetAllInputTypes()
}
