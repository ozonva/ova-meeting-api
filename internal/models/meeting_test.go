package models

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var m = Meeting{
	Id:     uuid.New(),
	UserId: 1,
	State: MeetingState{
		Id:   1,
		Name: "New",
	},
	Users: []uint64{},
}

func TestMeeting_ChangeState(t *testing.T) {
	assertions := assert.New(t)
	newState := MeetingState{
		Id:   2,
		Name: "Started",
	}
	m.ChangeState(newState)
	assertions.Equal(newState, m.State, "Should be equal")
}

func TestMeeting_GenerateId(t *testing.T) {
	assertions := assert.New(t)
	oldId := m.Id
	m.GenerateId()
	assertions.NotEqual(oldId, m.Id, "Should be not equal")
}

func TestMeeting_InviteUser(t *testing.T) {
	assertions := assert.New(t)
	m.InviteUser(2)
	assertions.Equal([]uint64{2}, m.Users, "Should be equal")
}

func TestMeeting_RemoveUser(t *testing.T) {
	assertions := assert.New(t)
	m.InviteUser(2)
	m.InviteUser(1)
	m.InviteUser(3)
	m.RemoveUser(2)
	assertions.Equal([]uint64{3, 1}, m.Users, "Should be equal")
}
