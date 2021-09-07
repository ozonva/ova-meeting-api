package api

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/ozonva/ova-meeting-api/internal/models"
	"github.com/ozonva/ova-meeting-api/internal/producer"
	desc "github.com/ozonva/ova-meeting-api/pkg/ova-meeting-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) UpdateMeetingV1(ctx context.Context, req *desc.UpdateMeetingRequestV1) (*empty.Empty, error) {

	log.Info().
		Caller().
		Str("ID", req.Id).
		Uint64("UserID", req.Userid).
		Str("Title", req.Title).
		Str("Date", req.Date.String()).
		Str("State", req.State.Name).
		Interface("Users", req.Users).Msg("")

	span := s.tracer.StartSpan("UpdateMeeting")
	defer span.Finish()

	meetingId, err := uuid.Parse(req.Id)

	if err != nil {
		log.Error().Err(err)
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}
	if err = req.Validate(); err != nil {
		log.Error().Err(err).Msg("Validation failed")
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}
	err = s.repo.UpdateMeeting(ctx, models.Meeting{
		ID:     meetingId,
		Title:  req.Title,
		UserID: req.Userid,
		Date:   req.Date.AsTime(),
		State: models.MeetingState{
			ID:   uint(req.State.Id),
			Name: req.State.Name,
		},
		Users: req.Users,
	})

	if err != nil {
		log.Error().Err(err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	s.reportEvent(producer.Update, meetingId)
	s.metrics.IncUpdate()

	return &empty.Empty{}, nil
}
