package domain

type KafkaConsumer interface {
	ConsumeMessages(topic string, messageHandler func([]byte) error) error
	Close() error
}
