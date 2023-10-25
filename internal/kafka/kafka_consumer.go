package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	reader *kafka.Reader
}

func NewKafkaConsumer(brokers []string, topic string, groupID string, dialer *kafka.Dialer) *KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
		Dialer:  dialer,
	})

	return &KafkaConsumer{
		reader: reader,
	}
}

func (kc *KafkaConsumer) ConsumeMessages(ctx context.Context, messageHandler func([]byte) error) error {
	defer kc.reader.Close()

	for {
		msg, err := kc.reader.ReadMessage(ctx)
		if err != nil {
			log.Println("Error reading message from kafka", err)
			return err
		}

		if err := messageHandler(msg.Value); err != nil {
			log.Println("Error handling message", err)
			return err
		}
	}
}

func (kc *KafkaConsumer) Close() error {
	return kc.reader.Close()
}
