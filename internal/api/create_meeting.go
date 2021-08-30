package api

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"

	desc "github.com/ozonva/ova-meeting-api/pkg/ova-meeting-api"
	"github.com/rs/zerolog/log"
)

func (s *Server) CreateMeetingV1(ctx context.Context, req *desc.AddMeetingRequestV1) (*empty.Empty, error) {

	log.Info().Msg("Create new meeting: " + req.String())

	return &empty.Empty{}, nil
}
