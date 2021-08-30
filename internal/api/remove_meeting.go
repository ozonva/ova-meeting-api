package api

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/ozonva/ova-meeting-api/pkg/ova-meeting-api"
	"github.com/rs/zerolog/log"
)

func (s *Server) RemoveMeetingV1(ctx context.Context, req *desc.MeetingIDRequestV1) (*empty.Empty, error) {

	log.Info().Msg("Delete Meeting with ID: " + req.GetId().String())

	return &empty.Empty{}, nil
}
