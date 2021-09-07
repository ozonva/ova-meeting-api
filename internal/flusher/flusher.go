package flusher

import (
	"context"
	"github.com/ozonva/ova-meeting-api/internal/models"
	"github.com/ozonva/ova-meeting-api/internal/repo"
	"github.com/ozonva/ova-meeting-api/internal/utils"
)

type flusher struct {
	chunkSize uint
	repo      repo.MeetingRepo
}

// Flusher interface to store data
type Flusher interface {
	Flush(meetings []models.Meeting) []models.Meeting
}

// NewFlusher return Flusher with batch save possibility
func NewFlusher(chunkSize uint, repo repo.MeetingRepo) Flusher {
	return &flusher{chunkSize: chunkSize, repo: repo}
}

// Flush flush Meetings using chunkSize per iteration. If error add Meetings to result
func (f *flusher) Flush(meetings []models.Meeting) []models.Meeting {
	var res []models.Meeting
	ctx := context.TODO()
	for _, meetingsPart := range utils.SplitSliceMeetings(meetings, f.chunkSize) {
		if _, err := f.repo.AddMeetings(ctx, meetingsPart); err != nil {
			res = append(res, meetingsPart...)
		}
	}
	return res
}
