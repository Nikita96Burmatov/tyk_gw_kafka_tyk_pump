package streams

import (
	"github.com/TykTechnologies/tyk/internal/bxais/helpers"
	"github.com/segmentio/kafka-go"
	"io"
	"net/http"
)

func (s *Stream) PrepareMessage(done <-chan int, ch <-chan http.Request) <-chan kafka.Message {
	brokerStream := make(chan kafka.Message)

	go func() {
		defer close(brokerStream)

		for {
			select {
			case <-done:
				return
			case req := <-ch:
				pathParts, _ := helpers.ParseURL(req.URL.Path)
				body, _ := io.ReadAll(req.Body)

				msg := s.broker.Message(
					pathParts[1],
					pathParts[2],
					body,
					s.broker.Headers(map[string]string{
						s.config.Kafka.Headers["recipient"]: pathParts[1],
						s.config.Kafka.Headers["path"]:      req.URL.Path,
					}),
				)

				go func() {
					brokerStream <- msg
				}()
			}
		}
	}()

	return brokerStream
}
