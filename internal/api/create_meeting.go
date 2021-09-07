package api

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/ozonva/ova-meeting-api/internal/models"
	"github.com/ozonva/ova-meeting-api/internal/producer"
	desc "github.com/ozonva/ova-meeting-api/pkg/ova-meeting-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateMeetingV1(ctx context.Context, req *desc.AddMeetingRequestV1) (*empty.Empty, error) {

	log.Info().
		Caller().
		Uint64("UserID", req.Userid).
		Str("Title", req.Title).
		Str("Date", req.Date.String()).
		Str("State", req.State.Name).
		Interface("Users", req.Users).
		Msg("Create meeting request")

	span := s.tracer.StartSpan("CreateMeeting")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("Validation failed")
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	uuids, err := s.repo.AddMeetings(ctx, []models.Meeting{{
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
		log.Error().Err(err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	if len(uuids) > 0 {
		s.reportEvent(producer.Create, uuids[0])
		s.metrics.IncCreate()
	}

	return &empty.Empty{}, nil
}
