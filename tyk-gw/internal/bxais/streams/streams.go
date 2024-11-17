package streams

import (
	"github.com/TykTechnologies/tyk/config"
	"github.com/TykTechnologies/tyk/internal/bxais/broker"
	"github.com/sirupsen/logrus"
)

type Stream struct {
	logger *logrus.Logger
	config config.Config
	broker *broker.Broker
}

func New(
	logger *logrus.Logger,
	config config.Config,
	broker *broker.Broker,
) *Stream {
	return &Stream{
		logger: logger,
		config: config,
		broker: broker,
	}
}
