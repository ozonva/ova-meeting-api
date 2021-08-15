package utils

import (
	"github.com/google/uuid"
	"github.com/ozonva/ova-meeting-api/internal/models"
)

// SliceToMap convert slice of Meeting to map
func SliceToMap(input []models.Meeting) map[uuid.UUID]models.Meeting {
	result := make(map[uuid.UUID]models.Meeting)

	for _, m := range input {
		result[m.ID] = m
	}

	return result
}
