package models

import (
	"errors"

	"github.com/google/uuid"
)

type Meeting struct {
	ID     uuid.UUID    `json:"id"`
	UserID uint64       `json:"user_id"`
	State  MeetingState `json:"state"`
	Users  []uint64     `json:"users,omitempty"`
}

// NewMeeting create new Meeting
func NewMeeting(userId uint64) Meeting {
	meeting := Meeting{
		UserID: userId,
		State: MeetingState{
			ID:   1,
			Name: "New",
		},
		Users: []uint64{},
	}
	meeting.GenerateId()
	return meeting
}

// String return string info about meeting
func (m Meeting) String() string {
	return "Meeting with id " + m.ID.String() + " has a " + m.State.Name + " state"
}

// GenerateId generate new meeting id
func (m *Meeting) GenerateId() {
	m.ID = uuid.New()
}

// InviteUser user with ID to the current meeting
func (m *Meeting) InviteUser(user uint64) {
	users := m.invitedUsersAsMap()
	if _, ok := users[user]; ok {
		return
	}
	m.Users = append(m.Users, user)
}

// RemoveUser remove user from meeting
func (m *Meeting) RemoveUser(user uint64) {
	curPos, err := m.userPos(user)
	if err == nil {
		m.Users[curPos] = m.Users[len(m.Users)-1]
		m.Users = m.Users[:len(m.Users)-1]
	}
}

// ChangeState change current meeting state to the new one
func (m *Meeting) ChangeState(newState MeetingState) {
	m.State = newState
}

// GetState return current meeting state
func (m Meeting) GetState() MeetingState {
	return m.State
}

func (m Meeting) invitedUsersAsMap() map[uint64]struct{} {
	result := make(map[uint64]struct{}, len(m.Users))
	for _, userId := range m.Users {
		if _, ok := result[userId]; ok {
			continue
		}
		result[userId] = struct{}{}
	}
	return result
}

func (m Meeting) userPos(user uint64) (int, error) {
	for p, v := range m.Users {
		if v == user {
			return p, nil
		}
	}
	return -1, errors.New("user not found")
}
