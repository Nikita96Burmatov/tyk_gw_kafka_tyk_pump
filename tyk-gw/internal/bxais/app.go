package bxais

import (
	"github.com/TykTechnologies/tyk/config"
	"github.com/TykTechnologies/tyk/internal/bxais/broker"
	"github.com/TykTechnologies/tyk/internal/bxais/streams"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	pipeline chan http.Request
	signals  chan os.Signal
	done     chan int
)

func HandleRequest(r http.Request) {
	pipeline <- r
}

func Run(
	logger *logrus.Logger,
	config config.Config,
) {
	pipeline = make(chan http.Request)
	signals = make(chan os.Signal, 1)
	done = make(chan int)
	defer close(done)
	defer close(signals)
	defer close(pipeline)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGKILL)

	b := broker.Kafka(config.Kafka.Addresses)
	s := streams.New(logger, config, b)

	// Producer pipeline
	messageStream := s.PrepareMessage(done, pipeline)
	s.WriteMessage(done, messageStream)

	// Consumer pipeline
	consumerStream := s.ReadMessages(done)
	s.DoRequest(done, consumerStream)

	for sig := range signals {
		logger.Error("BxAis app error: ", sig)
	}
}
