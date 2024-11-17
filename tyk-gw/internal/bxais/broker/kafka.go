package broker

import (
	"github.com/segmentio/kafka-go"
)

func Kafka(addresses []string) *Broker {
	return &Broker{
		Addresses: addresses,
	}
}

func (b *Broker) Writer() *kafka.Writer {
	return &kafka.Writer{
		Addr: kafka.TCP(b.Addresses...),
		// NOTE: When Topic is not defined here, each Message must define it instead.
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}
}

func (b *Broker) Reader(groupID, topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  b.Addresses,
		GroupID:  groupID,
		Topic:    topic,
		MaxBytes: 10e6, // 10MB
	})
}

func (b *Broker) Message(topic, key string, value []byte, header []kafka.Header) kafka.Message {
	return kafka.Message{
		Topic:   topic,
		Key:     []byte(key),
		Value:   value,
		Headers: header,
	}
}

func (b *Broker) Headers(headers map[string]string) []kafka.Header {
	//h := make([]kafka.Header, len(headers))
	var h []kafka.Header

	for key, value := range headers {
		h = append(h, kafka.Header{
			Key:   key,
			Value: []byte(value),
		})
	}

	return h
}
