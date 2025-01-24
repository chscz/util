package service

import "github.com/chscz/test_project/otel_sender/sender"

func (c *Container) initSenderTrace() *sender.TraceSender {
	c.initTraceOnce.Do(func() {
		c.traceSender = sender.NewTraceSender(c.config.Trace, c.logger)
	})
	return c.traceSender
}
