package api

import desc "github.com/ozonva/ova-meeting-api/pkg/ova-meeting-api"

type Server struct {
	desc.UnimplementedMeetingsServer
}

func NewApiServer() desc.MeetingsServer {
	return &Server{}
}
