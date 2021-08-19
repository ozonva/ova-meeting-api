package repo

import (
	"github.com/google/uuid"
	"github.com/ozonva/ova-meeting-api/internal/models"
)

// MeetingRepo storage interface of Meeting Entity
type MeetingRepo interface {
	AddMeetings(meetings []models.Meeting) error
	ListMeetings(limit, offset uint64) ([]models.Meeting, error)
	DescribeMeeting(meetingId uuid.UUID) (*models.Meeting, error)
}
