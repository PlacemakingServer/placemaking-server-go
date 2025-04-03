package repository

import (
	"encoding/json"
	"log"
	"placemaking-backend-go/config"
	"placemaking-backend-go/models"
)

func GetAllActivityTypes() ([]models.ActivityType, error){
	supabase := config.GetSupabase()

	response, _, err := supabase.From("activity_types").Select("*", "", false).Execute()
	
	if err != nil {
		log.Println("Error fetching activity types:", err)
		return nil, err
	}

	var activityTypes []models.ActivityType
	err = json.Unmarshal(response, &activityTypes)

	if err != nil {
		println("Error decoding activity types", err)
		return nil, err
	}

	return activityTypes, nil

}