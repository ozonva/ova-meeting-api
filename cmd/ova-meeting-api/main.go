package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/ozonva/ova-meeting-api/internal/api"
	desc "github.com/ozonva/ova-meeting-api/pkg/ova-meeting-api"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

const (
	readAllEntities    = 0
	grpcPort           = ":82"
	grpcServerEndpoint = "localhost:82"
	jsonEndpoint       = ":8081"
)

func runJSON() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterMeetingsHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(jsonEndpoint, mux)
	if err != nil {
		panic(err)
	}
}

func run() error {
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterMeetingsServer(s, api.NewApiServer())

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func main() {
	go runJSON()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
