package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/ozonva/ova-meeting-api/internal/api"
	"github.com/ozonva/ova-meeting-api/internal/config"
	"github.com/ozonva/ova-meeting-api/internal/metrics"
	"github.com/ozonva/ova-meeting-api/internal/producer"
	"github.com/ozonva/ova-meeting-api/internal/repo"
	desc "github.com/ozonva/ova-meeting-api/pkg/ova-meeting-api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	jaegerMetrics "github.com/uber/jaeger-lib/metrics"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

var cfg *config.Config

func regSignalHandler(ctx context.Context) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		defer signal.Stop(stop)
		<-stop
		log.Info().Msg("Stop signal received")
		cancel()
	}()
	return ctx
}

func run(ctx context.Context) error {
	mainRepo, err := getRepo()
	if err != nil {
		return err
	}
	mainProducer, err := producer.New(cfg.Broker.List)
	if err != nil {
		log.Error().Err(err).Msg("Producer: New")
		return err
	}
	defer mainProducer.Close()

	mainMetrics := metrics.New()

	tracer, tracingCloser, err := initTracing()
	if err != nil {
		log.Error().Err(err).Msg("Tracing: init")
		return err
	}
	defer func(tracingCloser io.Closer) {
		err := tracingCloser.Close()
		if err != nil {
			log.Fatal().Err(err).Msg("Tracing: close")
		}
	}(tracingCloser)

	errorGroup, ctx := errgroup.WithContext(ctx)
	errorGroup.Go(func() error { return runService(ctx, mainRepo, mainProducer, mainMetrics, tracer) })
	errorGroup.Go(func() error { return runGateway(ctx) })
	errorGroup.Go(func() error { return runMetrics(ctx) })

	return errorGroup.Wait()
}

func getRepo() (repo.MeetingRepo, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Database.User,
		cfg.Database.Pass,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		log.Error().Err(err).Msg("DB: Connect")
		return nil, err
	}

	meetingRepo := repo.NewRepo(db)
	return meetingRepo, nil
}

func runService(
	ctx context.Context,
	repo repo.MeetingRepo,
	producer producer.Producer,
	metrics metrics.Metrics,
	tracer opentracing.Tracer,
) error {
	grpcAddress := fmt.Sprintf("%s:%s", cfg.GRPC.Bind, cfg.GRPC.Port)
	listen, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Error().Err(err).Msg("GRPC: Listen")
		return err
	}

	srv := grpc.NewServer()
	desc.RegisterMeetingsServer(srv, api.NewApiServer(repo, producer, metrics, tracer, cfg.ChunkSize))

	srvErr := make(chan error)
	go func() {
		if err := srv.Serve(listen); err != nil {
			srvErr <- err
		}
	}()

	select {
	case err := <-srvErr:
		log.Error().Err(err).Msg("GRPC: Serve")
		return err

	case <-ctx.Done():
		srv.GracefulStop()
	}

	return nil
}

func runGateway(ctx context.Context) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	grpcEndpoint := fmt.Sprintf("%s:%s", cfg.GRPC.Bind, cfg.GRPC.Port)
	err := desc.RegisterMeetingsHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts)
	if err != nil {
		log.Error().Err(err).Msg("Gateway: Register API handler")
		return err
	}

	addr := fmt.Sprintf("%s:%s", cfg.Gateway.Bind, cfg.Gateway.Port)
	srv := &http.Server{Addr: addr, Handler: mux}

	srvErr := make(chan error)
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			srvErr <- err
		}
	}()

	select {
	case err := <-srvErr:
		log.Error().Err(err).Msg("Gateway: Serve")
		return err

	case <-ctx.Done():
		err := srv.Shutdown(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("Gateway: Shutdown")
			return err
		}
	}

	return nil
}

func runMetrics(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	addr := fmt.Sprintf("%s:%s", cfg.Metrics.Bind, cfg.Metrics.Port)
	srv := &http.Server{Addr: addr, Handler: mux}

	srvErr := make(chan error)
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			srvErr <- err
		}
	}()

	select {
	case err := <-srvErr:
		log.Error().Err(err).Msg("Metrics: Serve")
		return err

	case <-ctx.Done():
		err := srv.Shutdown(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("Metrics: Shutdown")
			return err
		}
	}

	return nil
}

func initTracing() (opentracing.Tracer, io.Closer, error) {
	localAgentAddr := fmt.Sprintf("%s:%s", cfg.Tracing.AgentHost, cfg.Tracing.AgentPort)

	cfg := jaegerConfig.Configuration{
		ServiceName: "ova-meeting-api",
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LocalAgentHostPort: localAgentAddr,
			LogSpans:           true,
		},
	}

	logger := jaeger.StdLogger
	metricsFactory := jaegerMetrics.NullFactory

	return cfg.NewTracer(
		jaegerConfig.Logger(logger),
		jaegerConfig.Metrics(metricsFactory),
	)
}

func getConfig() {
	cfg = config.Get()

	logLevels := map[string]zerolog.Level{
		"trace":    zerolog.TraceLevel,
		"debug":    zerolog.DebugLevel,
		"info":     zerolog.InfoLevel,
		"warn":     zerolog.WarnLevel,
		"error":    zerolog.ErrorLevel,
		"fatal":    zerolog.FatalLevel,
		"panic":    zerolog.PanicLevel,
		"disabled": zerolog.Disabled,
	}
	if level, ok := logLevels[cfg.LogLevel]; ok {
		zerolog.SetGlobalLevel(level)
	}
}

func main() {
	getConfig()
	ctx := regSignalHandler(context.Background())
	log.Info().Msg("Service: started")
	if err := run(ctx); err != nil {
		log.Fatal().Err(err).Msg("Service: stopped on error")
	}
	log.Info().Msg("Service: exited")
}
