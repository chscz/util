package service

import (
	"context"

	"github.com/chscz/test_project/otel_sender/config"
)

type Sender interface {
	Start(ctx context.Context) (int, int)
}

func (c *Container) InitSender(_ context.Context, senderType config.SenderType) Sender {
	switch senderType {
	case config.SenderTypeTrace:
		return c.initSenderTrace()
	case config.SenderTypeLog:
		return nil
	case config.SenderTypeMetric:
		return nil
	}
	return nil
}
