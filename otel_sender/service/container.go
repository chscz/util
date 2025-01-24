package service

import (
	"log/slog"
	"sync"

	"github.com/chscz/test_project/otel_sender/config"
	"github.com/chscz/test_project/otel_sender/sender"
)

type Container struct {
	logger         *slog.Logger
	initLoggerOnce sync.Once

	config config.Config

	traceSender   *sender.TraceSender
	initTraceOnce sync.Once
}

func NewContainer(cfg config.Config) *Container {
	return &Container{
		config: cfg,
	}
}

func (c *Container) Logger() *slog.Logger {
	c.initLoggerOnce.Do(func() {
		c.logger = slog.Default()
	})
	return c.logger
}

func (c *Container) Close() {
	return
}
