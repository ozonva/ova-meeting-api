package api_test

import (
	"context"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	"github.com/opentracing/opentracing-go/mocktracer"
	"github.com/ozonva/ova-meeting-api/internal/api"
	"github.com/ozonva/ova-meeting-api/internal/mocks"
	"github.com/ozonva/ova-meeting-api/internal/models"
	desc "github.com/ozonva/ova-meeting-api/pkg/ova-meeting-api"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Api", func() {
	var (
		mockCtrl     *gomock.Controller
		mockRepo     *mocks.MockMeetingRepo
		ctx          context.Context
		mockProducer *mocks.MockProducer
		mockMetrics  *mocks.MockMetrics
		mockTracer   *mocktracer.MockTracer
		meetings     = []models.Meeting{
			models.NewMeeting(1),
			models.NewMeeting(2),
			models.NewMeeting(3),
			models.NewMeeting(4),
			models.NewMeeting(5),
		}
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockMeetingRepo(mockCtrl)
		ctx = context.Background()
		mockProducer = mocks.NewMockProducer(mockCtrl)
		mockMetrics = mocks.NewMockMetrics(mockCtrl)
		mockTracer = mocktracer.New()
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Api create", func() {
		It("Should not error", func() {
			emptyUuid := uuid.UUID{}
			mockRepo.EXPECT().AddMeetings(ctx, gomock.Any()).Return([]uuid.UUID{emptyUuid}, nil).Times(5)
			mockProducer.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil).Times(5)
			mockMetrics.EXPECT().IncCreate().Times(5)
			apiServer := api.NewApiServer(mockRepo, mockProducer, mockMetrics, mockTracer, 2)
			for _, meeting := range meetings {
				_, err := apiServer.CreateMeetingV1(ctx, &desc.AddMeetingRequestV1{
					Userid: meeting.UserID,
					Title:  meeting.Title,
					Date: &timestamp.Timestamp{
						Seconds: meeting.Date.Unix(),
						Nanos:   int32(meeting.Date.Nanosecond()),
					},
					State: &desc.MeetingStateV1{
						Id:   uint64(meeting.State.ID),
						Name: meeting.State.Name,
					},
					Users: meeting.Users,
				})
				assert.Nil(GinkgoT(), err)
			}
		})
	})
	Context("Api multi create", func() {
		It("Should not error", func() {
			emptyUuid := uuid.UUID{}
			mockRepo.EXPECT().AddMeetings(ctx, gomock.Any()).Return([]uuid.UUID{emptyUuid, emptyUuid, emptyUuid, emptyUuid, emptyUuid}, nil).Times(1)
			mockProducer.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil).Times(5)
			mockMetrics.EXPECT().IncCreate().Times(5)

			meetingsRequest := make([]*desc.AddMeetingRequestV1, 0, 5)
			for _, meeting := range meetings {
				meetingsRequest = append(meetingsRequest, &desc.AddMeetingRequestV1{
					Userid: meeting.UserID,
					Title:  meeting.Title,
					Date: &timestamp.Timestamp{
						Seconds: meeting.Date.Unix(),
						Nanos:   int32(meeting.Date.Nanosecond()),
					},
					State: &desc.MeetingStateV1{
						Id:   uint64(meeting.State.ID),
						Name: meeting.State.Name,
					},
					Users: meeting.Users,
				})
			}
			apiServer := api.NewApiServer(mockRepo, mockProducer, mockMetrics, mockTracer, 5)
			_, err := apiServer.MultiCreateMeetingV1(ctx, &desc.MultiCreateMeetingRequestV1{
				Meetings: meetingsRequest,
			})
			assert.Nil(GinkgoT(), err)
		})
	})
	Context("Api describe", func() {
		It("Should not error", func() {
			mockRepo.EXPECT().DescribeMeeting(gomock.Any()).Return(models.Meeting{}, nil).Times(1)
			apiServer := api.NewApiServer(mockRepo, mockProducer, mockMetrics, mockTracer, 2)

			_, err := apiServer.DescribeMeetingV1(ctx, &desc.MeetingIDRequestV1{Id: meetings[0].ID.String()})
			assert.Nil(GinkgoT(), err)
		})
	})
	Context("Api list", func() {
		It("Should not error", func() {
			mockRepo.EXPECT().ListMeetings(gomock.Any(), gomock.Any()).Return([]models.Meeting{}, nil).Times(1)
			apiServer := api.NewApiServer(mockRepo, mockProducer, mockMetrics, mockTracer, 2)

			_, err := apiServer.ListMeetingsV1(ctx, &desc.ListMeetingsRequestV1{Limit: 10, Offset: 0})
			assert.Nil(GinkgoT(), err)
		})
	})
	Context("Api remove", func() {
		It("Should not error", func() {
			mockRepo.EXPECT().DeleteMeeting(ctx, gomock.Any()).Return(nil).Times(1)
			mockProducer.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			mockMetrics.EXPECT().IncDelete().Times(1)
			apiServer := api.NewApiServer(mockRepo, mockProducer, mockMetrics, mockTracer, 2)

			_, err := apiServer.RemoveMeetingV1(ctx, &desc.MeetingIDRequestV1{Id: meetings[0].ID.String()})
			assert.Nil(GinkgoT(), err)
		})
	})
	Context("Api update", func() {
		It("Should not error", func() {
			mockRepo.EXPECT().UpdateMeeting(ctx, gomock.Any()).Return(nil).Times(5)
			mockProducer.EXPECT().Send(gomock.Any(), gomock.Any()).Return(nil).Times(5)
			mockMetrics.EXPECT().IncUpdate().Times(5)
			apiServer := api.NewApiServer(mockRepo, mockProducer, mockMetrics, mockTracer, 2)
			for _, meeting := range meetings {
				_, err := apiServer.UpdateMeetingV1(ctx, &desc.UpdateMeetingRequestV1{
					Id:     meeting.ID.String(),
					Userid: meeting.UserID,
					Title:  meeting.Title,
					Date: &timestamp.Timestamp{
						Seconds: meeting.Date.Unix(),
						Nanos:   int32(meeting.Date.Nanosecond()),
					},
					State: &desc.MeetingStateV1{
						Id:   uint64(meeting.State.ID),
						Name: meeting.State.Name,
					},
					Users: meeting.Users,
				})
				assert.Nil(GinkgoT(), err)
			}
		})
	})
})
