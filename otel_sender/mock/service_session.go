package mock

import (
	"math/rand"
	"strings"

	"github.com/chscz/test_project/otel_sender/domain"
	"github.com/google/uuid"
)

var Services = []string{
	"react-sample-web@0.0.1",
	"io.imqa.examples.basic@1.0.0",
}

var ServiceSession []domain.ServiceSession

func GenerateServiceSession(count int) {
	ServiceSession = make([]domain.ServiceSession, count)
	for i := 0; i < count; i++ {
		service := strings.Split(Services[rand.Intn(len(Services))], "@")
		ServiceSession[i] = domain.ServiceSession{
			ServiceName:    service[0],
			ServiceVersion: service[1],
			ServiceKey:     "",
			SessionID:      uuid.New().String(),
		}
	}
}
