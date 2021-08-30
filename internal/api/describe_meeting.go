package api

import (
	"context"

	desc "github.com/ozonva/ova-meeting-api/pkg/ova-meeting-api"
	"github.com/rs/zerolog/log"
)

func (s *Server) DescribeMeetingV1(ctx context.Context, req *desc.MeetingIDRequestV1) (*desc.MeetingResponseV1, error) {

	log.Info().Msg("Describe Meeting with ID: " + req.GetId().String())

	return &desc.MeetingResponseV1{
		Id: req.Id,
	}, nil
}
