package saver_test

import (
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozonva/ova-meeting-api/internal/mocks"
	"github.com/ozonva/ova-meeting-api/internal/models"
	"github.com/ozonva/ova-meeting-api/internal/saver"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Saver", func() {
	var (
		mockCtrl    *gomock.Controller
		mockFlusher *mocks.MockFlusher
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
		mockFlusher = mocks.NewMockFlusher(mockCtrl)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("With cap 1", func() {
		It("Should save all", func() {
			mockFlusher.EXPECT().Flush(gomock.Any()).Times(5)
			sv := saver.NewSaver(1, mockFlusher, 1*time.Second)

			for _, m := range meetings {
				err := sv.Save(m)
				assert.Nil(GinkgoT(), err)
			}
			sv.Close()
		})
	})
	Context("With cap 2", func() {
		It("Should save all", func() {
			mockFlusher.EXPECT().Flush(gomock.Any()).Times(3)
			sv := saver.NewSaver(2, mockFlusher, 1*time.Second)

			for _, m := range meetings {
				err := sv.Save(m)
				assert.Nil(GinkgoT(), err)
			}
			sv.Close()
		})
	})
	Context("With cap 3", func() {
		It("Should save all", func() {
			mockFlusher.EXPECT().Flush(gomock.Any()).Times(2)
			sv := saver.NewSaver(3, mockFlusher, 1*time.Second)

			for _, m := range meetings {
				err := sv.Save(m)
				assert.Nil(GinkgoT(), err)
			}
			sv.Close()
		})
	})
	Context("With cap 4", func() {
		It("Should save all", func() {
			mockFlusher.EXPECT().Flush(gomock.Any()).Times(2)
			sv := saver.NewSaver(4, mockFlusher, 1*time.Second)

			for _, m := range meetings {
				err := sv.Save(m)
				assert.Nil(GinkgoT(), err)
			}
			sv.Close()
		})
	})
	Context("With cap 5", func() {
		It("Should save all", func() {
			mockFlusher.EXPECT().Flush(gomock.Any()).Times(1)
			sv := saver.NewSaver(5, mockFlusher, 1*time.Second)

			for _, m := range meetings {
				err := sv.Save(m)
				assert.Nil(GinkgoT(), err)
			}
			sv.Close()
		})
	})
})
