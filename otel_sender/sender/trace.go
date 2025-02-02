package sender

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/chscz/test_project/otel_sender/config"
	"github.com/chscz/test_project/otel_sender/domain"
	"google.golang.org/grpc"
)

type TraceSender struct {
	HTTPClient *http.Client
	GRPCClient *grpc.ClientConn
	logger     *slog.Logger
	timer      *time.Timer
	SendCount  int
	config     config.Trace
}

func NewTraceSender(cfg config.Trace, logger *slog.Logger) *TraceSender {
	sender := &TraceSender{
		config: cfg,
		logger: logger,
	}
	if cfg.ProtocolType == "http" {
		sender.HTTPClient = &http.Client{}
	} else if cfg.ProtocolType == "grpc" {
		//todo
	}

	sender.setTerminateConfig()

	return sender
}

func (t *TraceSender) Start(ctx context.Context) (success, failed int) {
	switch t.config.TerminationConfig.Mode {
	case config.ModeMaxCount:
		success, failed = t.countSender()
	case config.ModeDuration:
		success, failed = t.durationSender()
	case config.ModeEndDeadline:
		success, failed = t.deadlineSender()
	}
	return
}

func (t *TraceSender) countSender() (success, failed int) {
	for cnt := 0; cnt < t.config.TerminationConfig.MaxCount; cnt++ {
		data := domain.SetData()
		if err := t.send(data); err != nil {
			t.logger.Error(err.Error())
			failed++
		} else {
			success++
		}
	}
	return
}

func (t *TraceSender) durationSender() (success, failed int) {
	for {
		data := domain.SetData()
		if err := t.send(data); err != nil {
			t.logger.Error(err.Error())
			failed++
		} else {
			success++
		}
		select {
		case <-t.timer.C:
			return
		}
	}
}

func (t *TraceSender) deadlineSender() (success, failed int) {
	return
}

func (t *TraceSender) send(data io.Reader) error {
	req, err := http.NewRequest(http.MethodPost, t.config.URL, data)
	if err != nil {
		return err
	}
	if t.config.ContentType == config.ContentTypeJSON {
		req.Header.Add("Content-Type", "application/"+string(config.ContentTypeJSON))
	} else if t.config.ContentType == config.ContentTypeProtobuf {
		req.Header.Add("Content-Type", "application/x-"+string(config.ContentTypeProtobuf))
	}

	//resp, err := t.HTTPClient.Do(req)
	//if err != nil {
	//	return err
	//}
	//defer resp.Body.Close()
	//
	//b, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	return err
	//}
	//
	//fmt.Println(string(b))
	t.SendCount++
	return nil

}

func (t *TraceSender) setTerminateConfig() {
	switch t.config.TerminationConfig.Mode {
	case config.ModeMaxCount:
	case config.ModeDuration:
		t.timer = time.NewTimer(t.config.TerminationConfig.Duration)
		t.logger.Info("duration", t.config.TerminationConfig.Duration)
		go func() {
			<-t.timer.C
			t.timer.Stop()
			t.shutdown()
		}()
	case config.ModeEndDeadline:
		t.timer = time.NewTimer(time.Minute)
		t.logger.Info("deadline", t.config.TerminationConfig.EndDeadline)
		go func() {
			<-t.timer.C
			if t.config.TerminationConfig.EndDeadline == time.Now().Format(time.RFC3339) {
				t.shutdown()
			}
		}()
	default:
	}
	return
}

func (t *TraceSender) shutdown() {

}
