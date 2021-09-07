package api

import (
	"context"

	desc "github.com/ozonva/ova-meeting-api/pkg/ova-meeting-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ListMeetingsV1(ctx context.Context, req *desc.ListMeetingsRequestV1) (*desc.ListMeetingsResponseV1, error) {

	log.Info().Caller().Str("Params", req.String()).Msg("List meetings")

	span := s.tracer.StartSpan("ListMeetings")
	defer span.Finish()

	res, err := s.repo.ListMeetings(req.Limit, req.Offset)
	if err != nil {
		log.Error().Err(err)
		return nil, status.Error(codes.Internal, "internal error")
	}
	items := make([]*desc.MeetingResponseV1, 0, req.Limit)

	for _, m := range res {
		items = append(items, MeetingToResponseV1(m))
	}

	return &desc.ListMeetingsResponseV1{Items: items}, nil
}
