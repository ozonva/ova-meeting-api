package api

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	desc "github.com/ozonva/ova-meeting-api/pkg/ova-meeting-api"
	"github.com/rs/zerolog/log"
)

func (s *Server) RemoveMeetingV1(ctx context.Context, req *desc.MeetingIDRequestV1) (*empty.Empty, error) {

	log.Info().
		Caller().
		Msg("Delete Meeting with ID: " + req.GetId())

	meetingId, err := uuid.Parse(req.Id)

	if err != nil {
		log.Error().Err(err)
		return nil, err
	}
	err = s.repo.DeleteMeeting(ctx, meetingId)
	if err != nil {
		log.Error().Err(err)
		return nil, err
	}

	return &empty.Empty{}, nil
}
