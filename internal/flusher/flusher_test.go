package flusher_test

import (
	"fmt"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-meeting-api/internal/flusher"
	"github.com/ozonva/ova-meeting-api/internal/mocks"
	"github.com/ozonva/ova-meeting-api/internal/models"
)

var _ = Describe("Flusher", func() {
	var (
		mockCtrl *gomock.Controller
		mockRepo *mocks.MockMeetingRepo
	)
	const chunkSize = 2
	var (
		testFlusher flusher.Flusher
		meetings    = []models.Meeting{
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
		testFlusher = flusher.NewFlusher(chunkSize, mockRepo)
	})
	AfterEach(func() {
		mockCtrl.Finish()
	})
	Describe("Writing data to storage", func() {
		When("Write success", func() {
			AssertReturnNil := func(meetings []models.Meeting) {
				Expect(testFlusher.Flush(meetings)).To(BeNil())
			}
			Context("Write count less than chunkSize", func() {
				oneMeeting := meetings[:1]
				BeforeEach(func() {
					mockRepo.EXPECT().AddMeetings(oneMeeting).Return(nil).Times(1)
				})
				It("Should return nil", func() {
					AssertReturnNil(oneMeeting)
				})
			})
			Context("Write count equal chunkSize", func() {
				meetings := meetings[:chunkSize]
				BeforeEach(func() {
					mockRepo.EXPECT().AddMeetings(meetings).Return(nil).Times(1)
				})
				It("Should return nil", func() {
					AssertReturnNil(meetings)
				})
			})
			Context("Write count more than chunkSize", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddMeetings(meetings[:chunkSize]).Return(nil).Times(1),
						mockRepo.EXPECT().AddMeetings(meetings[chunkSize:chunkSize*2]).Return(nil).Times(1),
						mockRepo.EXPECT().AddMeetings(meetings[chunkSize*2:]).Return(nil).Times(1),
					)
				})
				It("Should return nil", func() {
					AssertReturnNil(meetings)
				})
			})
		})
		When("Write error", func() {
			err := fmt.Errorf("error writing data")
			AssertReturnMeetings := func(returnMeetings []models.Meeting) {
				Expect(testFlusher.Flush(meetings)).To(Equal(returnMeetings))
			}
			Context("All data", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddMeetings(meetings[:chunkSize]).Return(err).Times(1),
						mockRepo.EXPECT().AddMeetings(meetings[chunkSize:chunkSize*2]).Return(err).Times(1),
						mockRepo.EXPECT().AddMeetings(meetings[chunkSize*2:]).Return(err).Times(1),
					)
				})
				It("Should return all data", func() {
					AssertReturnMeetings(meetings)
				})
			})
			Context("Error write first chunk", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddMeetings(meetings[:chunkSize]).Return(err).Times(1),
						mockRepo.EXPECT().AddMeetings(meetings[chunkSize:chunkSize*2]).Return(nil).Times(1),
						mockRepo.EXPECT().AddMeetings(meetings[chunkSize*2:]).Return(nil).Times(1),
					)
				})
				It("Should return first chunk", func() {
					AssertReturnMeetings(meetings[:chunkSize])
				})
			})
		})
	})
})
