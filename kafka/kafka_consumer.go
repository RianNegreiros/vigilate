package kafka

import (
	"context"
	"log"

	"github.com/pusher/pusher-http-go"
	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	reader       *kafka.Reader
	pusherClient *pusher.Client
}

func NewKafkaConsumer(brokers []string, topic string, groupID string, dialer *kafka.Dialer, pusherClient *pusher.Client) *KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
		Dialer:  dialer,
	})

	return &KafkaConsumer{
		reader:       reader,
		pusherClient: pusherClient,
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

		messageContent := string(msg.Value)
		channelName := "kafka-messages"
		eventName := "kafka-message-received"
		data := map[string]string{"message": messageContent}

		if err := kc.pusherClient.Trigger(channelName, eventName, data); err != nil {
			log.Println("Error triggering pusher event", err)
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
