package kafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

type KafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer(producer *kafka.Producer) *KafkaProducer {
	return &KafkaProducer{producer}
}

func (kp *KafkaProducer) SendHealthCheckResultToKafka(result, topic string) {
	kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(result),
	}, nil)
}
