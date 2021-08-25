package saver

import (
	"errors"
	"sync"
	"time"

	"github.com/ozonva/ova-meeting-api/internal/flusher"
	"github.com/ozonva/ova-meeting-api/internal/models"
)

type Saver interface {
	Save(entity models.Meeting) error
	Close()
}

// NewSaver return Saver with periodic save possibility
func NewSaver(
	capacity uint,
	flusher flusher.Flusher,
	timeout time.Duration,
) Saver {
	saver := saver{
		flusher:    flusher,
		buffer:     make([]models.Meeting, 0, capacity),
		ticker:     time.NewTicker(timeout),
		signalChan: make(chan struct{}),
	}
	saver.periodicSave()
	return &saver
}

type saver struct {
	sync.Mutex
	flusher    flusher.Flusher
	buffer     []models.Meeting
	signalChan chan struct{}
	ticker     *time.Ticker
}

func (s *saver) Save(entity models.Meeting) error {
	if len(s.buffer) == cap(s.buffer) {
		s.flush()
	}
	return s.addToBuffer(entity)
}

func (s *saver) Close() {
	s.flush()
	close(s.signalChan)
}

func (s *saver) periodicSave() {
	go func(ch <-chan struct{}) {
		for {
			select {
			case <-s.ticker.C:
				s.flush()
			case _, ok := <-ch:
				if !ok {
					s.ticker.Stop()
					return
				}

			}
		}
	}(s.signalChan)
}

func (s *saver) addToBuffer(entity models.Meeting) error {
	s.Lock()
	defer s.Unlock()
	if len(s.buffer) == cap(s.buffer) {
		return errors.New("buffer is full")
	}
	s.buffer = append(s.buffer, entity)
	return nil
}

func (s *saver) flush() {
	s.Lock()
	defer s.Unlock()

	if len(s.buffer) == 0 {
		return
	}

	notSaved := s.flusher.Flush(s.buffer)
	s.buffer = s.buffer[:0]
	if len(notSaved) > 0 {
		// wait for next flush operation
		for _, val := range notSaved {
			s.buffer = append(s.buffer, val)
		}
	}
}
