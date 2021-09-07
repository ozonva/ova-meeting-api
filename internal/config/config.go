package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	ChunkSize uint   `yaml:"chunksize" env:"CHUNK_SIZE" env-default:"100"`
	LogLevel  string `yaml:"loglevel" env:"LOG_LEVEL"`

	GRPC struct {
		Bind string `yaml:"bind" env:"GRPC_BIND"`
		Port string `yaml:"port" env:"GRPC_PORT" env-default:"8082"`
	} `yaml:"grpc"`

	Gateway struct {
		Bind string `yaml:"bind" env:"GW_BIND"`
		Port string `yaml:"port" env:"GW_PORT" env-default:"8080"`
	} `yaml:"gateway"`

	Database struct {
		Host string `yaml:"host" env:"DB_HOST" env-default:"localhost"`
		Port string `yaml:"port" env:"DB_PORT" env-default:"5432"`
		User string `yaml:"user" env:"DB_USER" env-default:"postgres"`
		Pass string `yaml:"pass" env:"DB_PASS" env-default:"postgres"`
		Name string `yaml:"name" env:"DB_NAME" env-default:"postgres"`
	} `yaml:"database"`

	Metrics struct {
		Bind string `yaml:"bind" env:"METRICS_BIND" env-default:"localhost"`
		Port string `yaml:"port" env:"METRICS_PORT" env-default:"9100"`
	} `yaml:"metrics"`

	Broker struct {
		List []string `yaml:"list" env:"BROKERS" env-default:"localhost:9094"`
	} `yaml:"broker"`

	Tracing struct {
		AgentHost string `yaml:"agent_host" env:"TRACING_AGENT_HOST" env-default:"localhost"`
		AgentPort string `yaml:"agent_port" env:"TRACING_AGENT_PORT" env-default:"6831"`
	} `yaml:"tracing"`
}

var cfg *Config

func Get() *Config {
	cfg = &Config{}
	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err).Msg("No .env file found")
	}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Load configuration")
	}
	log.Info().Msg("Configuration loaded")
	return cfg
}
