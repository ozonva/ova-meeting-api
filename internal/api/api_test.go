package api_test

import (
	"context"

	"github.com/golang/mock/gomock"
	"github.com/golang/protobuf/ptypes/timestamp"
	. "github.com/onsi/ginkgo"
	"github.com/ozonva/ova-meeting-api/internal/api"
	"github.com/ozonva/ova-meeting-api/internal/mocks"
	"github.com/ozonva/ova-meeting-api/internal/models"
	desc "github.com/ozonva/ova-meeting-api/pkg/ova-meeting-api"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Api", func() {
	var (
		mockCtrl *gomock.Controller
		mockRepo *mocks.MockMeetingRepo
		ctx      context.Context
		meetings = []models.Meeting{
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
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Api create", func() {
		It("Should not error", func() {
			mockRepo.EXPECT().AddMeetings(gomock.Any()).Return(nil).Times(5)
			apiServer := api.NewApiServer(mockRepo)
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
	Context("Api describe", func() {
		It("Should not error", func() {
			mockRepo.EXPECT().DescribeMeeting(gomock.Any()).Return(models.Meeting{}, nil).Times(1)
			apiServer := api.NewApiServer(mockRepo)

			_, err := apiServer.DescribeMeetingV1(ctx, &desc.MeetingIDRequestV1{Id: meetings[0].ID.String()})
			assert.Nil(GinkgoT(), err)
		})
	})
	Context("Api list", func() {
		It("Should not error", func() {
			mockRepo.EXPECT().ListMeetings(gomock.Any(), gomock.Any()).Return([]models.Meeting{}, nil).Times(1)
			apiServer := api.NewApiServer(mockRepo)

			_, err := apiServer.ListMeetingsV1(ctx, &desc.ListMeetingsRequestV1{Limit: 10, Offset: 0})
			assert.Nil(GinkgoT(), err)
		})
	})
	Context("Api remove", func() {
		It("Should not error", func() {
			mockRepo.EXPECT().DeleteMeeting(gomock.Any()).Return(nil).Times(1)
			apiServer := api.NewApiServer(mockRepo)

			_, err := apiServer.RemoveMeetingV1(ctx, &desc.MeetingIDRequestV1{Id: meetings[0].ID.String()})
			assert.Nil(GinkgoT(), err)
		})
	})
	Context("Api update", func() {
		It("Should not error", func() {
			mockRepo.EXPECT().UpdateMeeting(gomock.Any()).Return(nil).Times(5)
			apiServer := api.NewApiServer(mockRepo)
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
