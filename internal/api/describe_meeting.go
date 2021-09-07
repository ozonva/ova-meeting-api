package api

import (
	"context"

	"github.com/google/uuid"
	desc "github.com/ozonva/ova-meeting-api/pkg/ova-meeting-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) DescribeMeetingV1(ctx context.Context, req *desc.MeetingIDRequestV1) (*desc.MeetingResponseV1, error) {

	log.Info().
		Caller().
		Msg("Describe Meeting with ID: " + req.Id)

	span := s.tracer.StartSpan("DescribeMeeting")
	defer span.Finish()

	reqUuid, err := uuid.Parse(req.Id)

	if err != nil {
		log.Error().Err(err)
		return nil, err
	}
	m, err := s.repo.DescribeMeeting(reqUuid)
	if err != nil {
		log.Error().Err(err)
		return nil, status.Error(codes.Internal, "internal error")
	}
	return MeetingToResponseV1(m), nil
}
