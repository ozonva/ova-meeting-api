package models

type MeetingState struct {
	ID   uint   `json:"id" db:"state_id"`
	Name string `json:"name" db:"state_name"`
}
