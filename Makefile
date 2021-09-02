include .env
LOCAL_BIN:=$(CURDIR)/bin
GOBIN?=$(GOPATH)/bin
DBSTRING:="postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(DB_HOST):$(DB_LOCAL_PORT)/$(POSTGRES_DB)?sslmode=disable"

.PHONY: run, test, generate, deps, migrate
export GO111MODULE=on
export GOPROXY=https://proxy.golang.org|direct


.PHONY: build
build:
	go build -o $(CURDIR)/bin/project $(CURDIR)/cmd/main.go

.PHONY: run
run:
	go run cmd/ova-meeting-api/main.go

.PHONY: test
test:
	go test ./internal/utils/
	go test ./internal/models/
	go test ./internal/flusher/
	go test ./internal/saver/
	go test ./internal/api/

.PHONY: generate
generate:
	go generate ./...
	PATH="${PATH}:$(LOCAL_BIN)" GOBIN=$(LOCAL_BIN) protoc -I $(CURDIR)/proto \
        --go_out=$(CURDIR)/pkg --go_opt=paths=source_relative \
        --go-grpc_out=$(CURDIR)/pkg --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=$(CURDIR)/pkg --grpc-gateway_opt paths=source_relative \
        $(CURDIR)/proto/ova-meeting-api/*.proto

.PHONY: deps
deps: install-go-deps

.PHONY: .install-go-deps
install-go-deps:
	ls go.mod || go mod init github.com/ozonva/ova-meeting-api
	GOBIN=$(LOCAL_BIN) go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	GOBIN=$(LOCAL_BIN) go get -u github.com/golang/protobuf/proto
	GOBIN=$(LOCAL_BIN) go get -u github.com/golang/protobuf/protoc-gen-go
	GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc
	GOBIN=$(LOCAL_BIN) go get -u github.com/grpc-ecosystem/grpc-gateway
	GOBIN=$(LOCAL_BIN) go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	GOBIN=$(LOCAL_BIN) go get -u github.com/onsi/ginkgo
	GOBIN=$(LOCAL_BIN) go get -u github.com/onsi/gomega
	GOBIN=$(LOCAL_BIN) go get -u github.com/golang/mock
	GOBIN=$(LOCAL_BIN) go get -u github.com/rs/zerolog/log
	GOBIN=$(LOCAL_BIN) go get -u github.com/pressly/goose/v3/cmd/goose

.PHONY: migrate
migrate:
	PATH="${PATH}:$(LOCAL_BIN)" GOBIN=$(LOCAL_BIN) GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DBSTRING) goose -dir db/migrations status
	PATH="${PATH}:$(LOCAL_BIN)" GOBIN=$(LOCAL_BIN) GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DBSTRING) goose -dir db/migrations up