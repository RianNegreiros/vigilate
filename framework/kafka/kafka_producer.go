package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	writer *kafka.Writer
}

func NewKafkaProducer(brokers []string, topic string, dialer *kafka.Dialer) *KafkaProducer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
		Dialer:  dialer,
	})

	return &KafkaProducer{
		writer: writer,
	}
}

func (kp *KafkaProducer) SendHealthCheckResultToKafka(result, topic string) error {
	message := kafka.Message{
		Key:   []byte("health-check-result"),
		Value: []byte(result),
	}

	err := kp.writer.WriteMessages(context.Background(), message)
	if err != nil {
		log.Println("Error sending health check result to Kafka: ", err)
		return err
	}

	return nil
}

func (kp *KafkaProducer) Close() {
	kp.writer.Close()
}
