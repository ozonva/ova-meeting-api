package api

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	desc "github.com/ozonva/ova-meeting-api/pkg/ova-meeting-api"
	"github.com/rs/zerolog/log"
)

func (s *Server) ListMeetingsV1(ctx context.Context, req *empty.Empty) (*desc.ListMeetingsResponseV1, error) {

	log.Info().Msg("List all meetings")

	return &desc.ListMeetingsResponseV1{
		Items: []*desc.MeetingResponseV1{},
	}, nil
}
