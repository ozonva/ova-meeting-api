package api

import (
	"context"
	"github.com/ozonva/ova-meeting-api/internal/producer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	desc "github.com/ozonva/ova-meeting-api/pkg/ova-meeting-api"
	"github.com/rs/zerolog/log"
)

func (s *Server) RemoveMeetingV1(ctx context.Context, req *desc.MeetingIDRequestV1) (*empty.Empty, error) {

	log.Info().
		Caller().
		Msg("Delete Meeting with ID: " + req.GetId())

	span := s.tracer.StartSpan("RemoveMeeting")
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
	err = s.repo.DeleteMeeting(ctx, meetingId)
	if err != nil {
		log.Error().Err(err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	s.reportEvent(producer.Delete, meetingId)
	s.metrics.IncDelete()

	return &empty.Empty{}, nil
}
