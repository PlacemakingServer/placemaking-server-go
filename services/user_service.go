package services

import (
    "placemaking-backend-go/models"
    repository "placemaking-backend-go/repositories"
)

func FetchUserById(id string) (models.User, error) {
    user, err := repository.GetUserById(id)
		if err != nil {
			return models.User{}, err
		}
		return user, nil 
}
