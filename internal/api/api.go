package api

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/ozonva/ova-meeting-api/internal/models"
	"github.com/ozonva/ova-meeting-api/internal/repo"
	desc "github.com/ozonva/ova-meeting-api/pkg/ova-meeting-api"
)

type Server struct {
	desc.UnimplementedMeetingsServer
	repo repo.MeetingRepo
}

func NewApiServer(r repo.MeetingRepo) desc.MeetingsServer {
	return &Server{repo: r}
}

func MeetingToResponseV1(m models.Meeting) *desc.MeetingResponseV1 {
	return &desc.MeetingResponseV1{
		Id: &desc.UUID{
			Value: m.ID.String(),
		},
		Userid: m.UserID,
		Title:  m.Title,
		Date: &timestamp.Timestamp{
			Seconds: m.Date.Unix(),
			Nanos:   int32(m.Date.Nanosecond()),
		},
		State: &desc.MeetingStateV1{
			Id:   uint64(m.State.ID),
			Name: m.State.Name,
		},
		Users: m.Users,
	}
}
