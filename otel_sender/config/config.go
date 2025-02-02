package config

import (
	"encoding/json"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type SenderType string

const (
	SenderTypeTrace  SenderType = "trace"
	SenderTypeLog    SenderType = "log"
	SenderTypeMetric SenderType = "metric"
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
	EndDeadline string          `env:"END_DEADLINE"`
}

type TerminationMode int8

const (
	ModeMaxCount TerminationMode = iota
	ModeDuration
	ModeEndDeadline
)

const (
	traceURL  = "http://localhost:8000/v1/traces"
	logURL    = "http://localhost:8000/v1/logs"
	metricURL = "http://localhost:8000/v1/metrics"
)

type Config struct {
	Trace        Trace  `envPrefix:"TRACE_"`
	Log          Log    `envPrefix:"LOG_"`
	Metric       Metric `envPrefix:"METRIC_"`
	SessionCount int    `env:"SESSION_COUNT"`
}

func (c Config) String() string {
	b, _ := json.MarshalIndent(c, "", "  ")
	return string(b)
}

func LoadFromEnv() (Config, error) {
	cfg := Config{
		Trace: Trace{
			URL:          "",
			Enable:       false,
			ContentType:  "",
			ProtocolType: "",
			TerminationConfig: TerminationConfig{
				Mode:        0,
				MaxCount:    0,
				Duration:    0,
				EndDeadline: "",
			},
		},
		Log:          Log{},
		Metric:       Metric{},
		SessionCount: 10,
	}

	if err := godotenv.Load(".env"); err != nil {
		return Config{}, err
	}

	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (c Config) GetSenderList() []SenderType {
	var senders []SenderType
	if c.Trace.Enable {
		senders = append(senders, SenderTypeTrace)
	}
	if c.Log.Enable {
		senders = append(senders, SenderTypeLog)
	}
	if c.Metric.Enable {
		senders = append(senders, SenderTypeMetric)
	}
	return senders
}
