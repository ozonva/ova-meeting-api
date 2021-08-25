package internal

//go:generate mockgen -destination=./mocks/flusher_mock.go -package=mocks github.com/ozonva/ova-meeting-api/internal/flusher Flusher

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozonva/ova-meeting-api/internal/repo MeetingRepo

//go:generate mockgen -destination=./mocks/saver_mock.go -package=mocks github.com/ozonva/ova-meeting-api/internal/saver Saver
