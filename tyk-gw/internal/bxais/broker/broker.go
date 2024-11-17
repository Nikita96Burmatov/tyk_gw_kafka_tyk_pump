package broker

type Broker struct {
	Addresses []string
}

//
//type Broker[Writer any, Reader any, Message any, Header any] interface {
//	Writer() *Writer
//	Reader()
//	Message(topic, key string, value []byte, header []Header) Message
//	Header(key string, value string) []Header
//}
//
//func Kafka(address string) Broker[kafka.Writer, kafka.Reader, kafka.Message, kafka.Header] {
//	return New(address)
//}
