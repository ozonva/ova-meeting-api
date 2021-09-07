package api

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	tracerlog "github.com/opentracing/opentracing-go/log"
	"github.com/ozonva/ova-meeting-api/internal/models"
	"github.com/ozonva/ova-meeting-api/internal/producer"
	"github.com/ozonva/ova-meeting-api/internal/utils"
	desc "github.com/ozonva/ova-meeting-api/pkg/ova-meeting-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) MultiCreateMeetingV1(ctx context.Context, req *desc.MultiCreateMeetingRequestV1) (*empty.Empty, error) {

	log.Info().
		Caller().
		Int("num_items", len(req.GetMeetings())).
		Msg("Multi create meetings request")

	span := s.tracer.StartSpan("MultiCreateMeeting")
	defer span.Finish()

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("Validation failed")
		return nil, status.Error(codes.InvalidArgument, "invalid argument")
	}

	reqMeetings := req.GetMeetings()
	meetings := make([]models.Meeting, 0, len(reqMeetings))
	for _, m := range reqMeetings {
		meeting := models.Meeting{
			Title:  m.Title,
			UserID: m.Userid,
			Date:   m.Date.AsTime(),
			State: models.MeetingState{
				ID:   uint(m.State.Id),
				Name: m.State.Name,
			},
			Users: m.Users,
		}
		meetings = append(meetings, meeting)
	}

	uuids := make([]uuid.UUID, 0, len(meetings))
	chunks := utils.SplitSliceMeetings(meetings, s.chunkSize)

	addChunk := func(ctx context.Context, meetings []models.Meeting) ([]uuid.UUID, error) {
		chunkSpan := s.tracer.StartSpan("Chunk", opentracing.ChildOf(span.Context()))
		chunkSpan.LogFields(tracerlog.Int("num_items", len(meetings)))
		defer chunkSpan.Finish()
		return s.repo.AddMeetings(ctx, meetings)
	}

	for chunkId, chunk := range chunks {
		newUuids, err := addChunk(ctx, chunk)
		if err != nil {
			log.Error().
				Int("chunk", chunkId).
				Err(err).
				Msg("Repo: add chunk failed")

			return &empty.Empty{}, status.Error(codes.Internal, "internal error")
		}
		uuids = append(uuids, newUuids...)

		for _, mUuid := range newUuids {
			s.reportEvent(producer.Create, mUuid)
			s.metrics.IncCreate()
		}
	}

	return &empty.Empty{}, nil
}
