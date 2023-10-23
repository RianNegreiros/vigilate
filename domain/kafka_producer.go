package domain

type KafkaProducer interface {
	SendHealthCheckResultToKafka(result, topic string) error
}
