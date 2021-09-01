package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/ozonva/ova-meeting-api/internal/api"
	"github.com/ozonva/ova-meeting-api/internal/connection"
	"github.com/ozonva/ova-meeting-api/internal/repo"
	desc "github.com/ozonva/ova-meeting-api/pkg/ova-meeting-api"
	"google.golang.org/grpc"
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

func run(dbConn *sqlx.DB) error {
	ctx := context.TODO()
	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterMeetingsServer(s, api.NewApiServer(repo.NewRepo(dbConn, ctx)))

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	go runJSON()

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_LOCAL_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	dbConn := connection.Connect(dsn)
	if err := dbConn.Ping(); err != nil {
		log.Fatal(err)
	}
	if err := run(dbConn); err != nil {
		log.Fatal(err)
	}
}
