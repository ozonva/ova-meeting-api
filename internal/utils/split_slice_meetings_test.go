package utils

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ozonva/ova-meeting-api/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestSplitSliceInterface(t *testing.T) {
	assertions := assert.New(t)
	u1 := uuid.New()
	u2 := uuid.New()
	u3 := uuid.New()
	u4 := uuid.New()

	m1 := models.Meeting{
		ID:     u1,
		UserID: 1,
		State: models.MeetingState{
			ID:   1,
			Name: "New",
		},
		Users: []uint64{1, 2},
	}
	m2 := models.Meeting{
		ID:     u2,
		UserID: 2,
		State: models.MeetingState{
			ID:   2,
			Name: "Cancelled",
		},
		Users: []uint64{2, 3},
	}
	m3 := models.Meeting{
		ID:     u3,
		UserID: 3,
		State: models.MeetingState{
			ID:   3,
			Name: "NewNew",
		},
		Users: []uint64{1, 2, 3, 4},
	}
	m4 := models.Meeting{
		ID:     u4,
		UserID: 4,
		State: models.MeetingState{
			ID:   4,
			Name: "State",
		},
		Users: []uint64{2, 3, 1, 2},
	}

	var testParams = []struct {
		testName string
		input    []models.Meeting
		chunk    int
		output   [][]models.Meeting
	}{
		{"Empty splice", []models.Meeting{}, 1, [][]models.Meeting{}},
		{"Split chunk size 1", []models.Meeting{m1, m2, m3, m4}, 1, [][]models.Meeting{{m1}, {m2}, {m3}, {m4}}},
		{"Split chunk size 2", []models.Meeting{m1, m2, m3, m4}, 2, [][]models.Meeting{{m1, m2}, {m3, m4}}},
		{"Split chunk size 3", []models.Meeting{m1, m2, m3, m4}, 3, [][]models.Meeting{{m1, m2, m3}, {m4}}},
		{"Split chunk size 4", []models.Meeting{m1, m2, m3, m4}, 4, [][]models.Meeting{{m1, m2, m3, m4}}},
	}
	for _, testParam := range testParams {
		assertions.Equal(testParam.output, SplitSliceMeetings(testParam.input, uint(testParam.chunk)), "Should be equal. "+testParam.testName)
	}

	assert.Panics(t, func() { SplitSliceMeetings([]models.Meeting{m1, m2}, 0) })
}
