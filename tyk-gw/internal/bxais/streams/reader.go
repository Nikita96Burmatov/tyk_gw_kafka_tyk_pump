package streams

import (
	"bytes"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"net/http"
	"net/url"
	"strings"
)

func (s *Stream) ReadMessages(done <-chan int) <-chan http.Request {
	stream := make(chan http.Request)

	listenPaths := s.config.Kafka.ListenPaths

	for _, p := range listenPaths {
		go func(path string) {
			u, _ := url.Parse(path)
			pathParts := strings.Split(u.Path, "/")

			r := s.broker.Reader(path, pathParts[1])
			defer func(r *kafka.Reader) {
				err := r.Close()
				if err != nil {
					s.logger.Error(err)
				}
			}(r)

			for {
				select {
				case <-done:
					return
				default:
					ctx := context.Background()

					// Получаем последнее доступное сообщение
					m, err := r.FetchMessage(ctx)
					if err != nil {
						s.logger.Error("failed to fetch kafka messages: ", err)
						continue
					}

					var url string

					if s.config.ListenAddress == "" {
						url = fmt.Sprintf("http://localhost:%d", s.config.ListenPort)
					} else {
						url = fmt.Sprintf("%s:%d", s.config.ListenAddress, s.config.ListenPort)
					}

					for _, header := range m.Headers {
						if header.Key == s.config.Kafka.Headers["path"] {
							url += string(header.Value)
						}
					}

					req, err := http.NewRequest("POST", url, bytes.NewBuffer(m.Value))
					if err != nil {
						s.logger.Error("failed to make request body: ", err)
						continue
					}

					for _, h := range m.Headers {
						if h.Key != "" && string(h.Value) != "" {
							req.Header.Set(h.Key, string(h.Value))
						}
					}

					// Отправляем запрос
					//res, err := helpers.Fetch(req)
					//if err != nil {
					//	return
					//}

					// Если ошибок нет, говорим брокеру что сообщение прочитано
					if err := r.CommitMessages(ctx, m); err != nil {
						s.logger.Error("failed to commit kafka messages: ", err)
					}

					stream <- *req
				}
			}

			//for {
			//
			//}
		}(p)
	}

	return stream
}
