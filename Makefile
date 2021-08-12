build:
	go build -o bin/main cmd/ova-meeting-api/main.go

run:
	go run cmd/ova-meeting-api/main.go

test:
	 go test ./internal/utils/
	 go test ./internal/models/