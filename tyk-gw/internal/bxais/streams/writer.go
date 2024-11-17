package streams

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

func (s *Stream) WriteMessage(done <-chan int, ch <-chan kafka.Message) {
	go func() {
		w := s.broker.Writer()

		for {
			select {
			case <-done:
				return
			case msg := <-ch:
				const retries = 3

				for i := 0; i < retries; i++ {
					err := w.WriteMessages(context.Background(), msg)

					if err != nil {
						s.logger.Error("failed to write message: ", err)
						time.Sleep(time.Millisecond * 250)
						continue
					}

					break
				}
			}
		}
	}()
}
