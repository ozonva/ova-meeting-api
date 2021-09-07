package api

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/ozonva/ova-meeting-api/internal/metrics"
	"github.com/ozonva/ova-meeting-api/internal/models"
	"github.com/ozonva/ova-meeting-api/internal/producer"
	"github.com/ozonva/ova-meeting-api/internal/repo"
	desc "github.com/ozonva/ova-meeting-api/pkg/ova-meeting-api"
	"github.com/rs/zerolog/log"
)

type Server struct {
	desc.UnimplementedMeetingsServer
	repo      repo.MeetingRepo
	producer  producer.Producer
	metrics   metrics.Metrics
	tracer    opentracing.Tracer
	chunkSize uint
}

func NewApiServer(r repo.MeetingRepo,
	producer producer.Producer,
	metrics metrics.Metrics,
	tracer opentracing.Tracer,
	chunkSize uint,
) desc.MeetingsServer {
	return &Server{repo: r, producer: producer, metrics: metrics, tracer: tracer, chunkSize: chunkSize}
}

const (
	producerTopic = "ova-meeting-api_events"
)

func (s *Server) reportEvent(typ producer.EventType, meetingId uuid.UUID) {
	ev := producer.PrepareEvent(typ, meetingId)
	err := s.producer.Send(producerTopic, ev)
	if err != nil {
		log.Error().Err(err).Msg("Producer: Send event")
	}
}

func MeetingToResponseV1(m models.Meeting) *desc.MeetingResponseV1 {
	return &desc.MeetingResponseV1{
		Id: &desc.UUID{
			Value: m.ID.String(),
		},
		Userid: m.UserID,
		Title:  m.Title,
		Date: &timestamp.Timestamp{
			Seconds: m.Date.Unix(),
			Nanos:   int32(m.Date.Nanosecond()),
		},
		State: &desc.MeetingStateV1{
			Id:   uint64(m.State.ID),
			Name: m.State.Name,
		},
		Users: m.Users,
	}
}
