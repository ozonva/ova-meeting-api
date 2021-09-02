package api

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/ozonva/ova-meeting-api/internal/models"
	desc "github.com/ozonva/ova-meeting-api/pkg/ova-meeting-api"
	"github.com/rs/zerolog/log"
)

func (s *Server) CreateMeetingV1(ctx context.Context, req *desc.AddMeetingRequestV1) (*empty.Empty, error) {

	log.Info().
		Caller().
		Uint64("UserID", req.Userid).
		Str("Title", req.Title).
		Str("Date", req.Date.String()).
		Str("State", req.State.Name).
		Interface("Users", req.Users).Msg("")

	err := s.repo.AddMeetings(ctx, []models.Meeting{{
		Title:  req.Title,
		UserID: req.Userid,
		Date:   req.Date.AsTime(),
		State: models.MeetingState{
			ID:   uint(req.State.Id),
			Name: req.State.Name,
		},
		Users: req.Users,
	}})

	if err != nil {
		log.Err(err)
	}

	return &empty.Empty{}, nil
}
