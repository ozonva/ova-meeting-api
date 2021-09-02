package api

import (
	"context"

	"github.com/google/uuid"
	desc "github.com/ozonva/ova-meeting-api/pkg/ova-meeting-api"
	"github.com/rs/zerolog/log"
)

func (s *Server) DescribeMeetingV1(ctx context.Context, req *desc.MeetingIDRequestV1) (*desc.MeetingResponseV1, error) {

	log.Info().Msg("Describe Meeting with ID: " + req.Id)
	reqUuid, err := uuid.Parse(req.Id)
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}
	m, err := s.repo.DescribeMeeting(reqUuid)
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}
	return MeetingToResponseV1(m), nil
}
