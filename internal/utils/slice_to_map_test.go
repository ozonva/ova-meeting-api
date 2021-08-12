package utils

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ozonva/ova-meeting-api/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestSliceToMap(t *testing.T) {
	assertions := assert.New(t)
	u1 := uuid.New()
	u2 := uuid.New()
	var meetings = []models.Meeting{
		{
			Id:     u1,
			UserId: 1,
			State: models.MeetingState{
				Id:   1,
				Name: "New",
			},
			Users: []uint64{1, 2},
		},
		{
			Id:     u2,
			UserId: 2,
			State: models.MeetingState{
				Id:   2,
				Name: "Cancelled",
			},
			Users: []uint64{2, 3},
		},
	}
	assertions.Equal(map[uuid.UUID]models.Meeting{u1: meetings[0], u2: meetings[1]}, SliceToMap(meetings), "Should be equal")
}
