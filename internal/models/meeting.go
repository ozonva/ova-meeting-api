package models

import (
	"github.com/google/uuid"
)

type Meeting struct {
	Id     uuid.UUID    `json:"id"`
	UserId uint64       `json:"user_id"`
	State  MeetingState `json:"state"`
	Users  []uint64     `json:"users,omitempty"`
}

// String return string info about meeting
func (m Meeting) String() string {
	return "Meeting with id " + m.Id.String() + " has a " + m.State.Name + " state"
}

// GenerateId generate new meeting id
func (m *Meeting) GenerateId() {
	m.Id = uuid.New()
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
	curPos := m.userPos(user)
	if curPos >= 0 {
		m.Users[curPos] = m.Users[len(m.Users)-1]
		m.Users = m.Users[:len(m.Users)-1]
	}
}

// ChangeState change current meeting state to the new one
func (m *Meeting) ChangeState(newState MeetingState) {
	m.State = newState
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

func (m Meeting) userPos(user uint64) int {
	for p, v := range m.Users {
		if v == user {
			return p
		}
	}
	return -1
}
