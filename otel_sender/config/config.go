package config

import (
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type ContentType string

const (
	ContentTypeJSON     ContentType = "json"
	ContentTypeProtobuf ContentType = "protobuf"
)

type ProtocolType string

const (
	ProtocolTypeHTTP ProtocolType = "http"
	ProtocolTypeGRPC ProtocolType = "grpc"
)

type TerminationConfig struct {
	Mode        TerminationMode `env:"MODE"`
	MaxCount    int             `env:"MAX_COUNT"`
	Duration    time.Duration   `env:"DURATION"`
	EndDeadline time.Time       `env:"END_DEADLINE"`
}

type TerminationMode int8

const (
	ModeMaxCount TerminationMode = iota
	ModeDuration
	ModeEndDeadline
)

type Config struct {
	Trace  Trace  `envPrefix:"TRACE_"`
	Log    Log    `envPrefix:"LOG_"`
	Metric Metric `envPrefix:"METRIC_"`
}

func LoadFromEnv() (Config, error) {
	cfg := Config{
		Trace: Trace{
			Enable:       false,
			ContentType:  "",
			ProtocolType: "",
			TerminationConfig: TerminationConfig{
				Mode:        0,
				MaxCount:    0,
				Duration:    0,
				EndDeadline: time.Time{},
			},
		},
		Log:    Log{},
		Metric: Metric{},
	}

	if err := godotenv.Load(".env"); err != nil {
		return Config{}, err
	}

	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
