package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pusher/pusher-http-go"
)

type KafkaConsumer struct {
	consumer     *kafka.Consumer
	pusherClient *pusher.Client
}

func NewKafkaConsumer(consumer *kafka.Consumer, pusherClient *pusher.Client) *KafkaConsumer {
	return &KafkaConsumer{
		consumer:     consumer,
		pusherClient: pusherClient,
	}
}

func (kc *KafkaConsumer) ConsumeMessages(topic string, messageHandler func([]byte) error) error {
	kc.consumer.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := kc.consumer.ReadMessage(-1)
		if err == nil {
			messageContent := string(msg.Value)

			channelName := "kafka-messages"
			eventName := "kafka-message-received"
			data := map[string]string{"message": messageContent}
			err := kc.pusherClient.Trigger(channelName, eventName, data)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
}

func (kc *KafkaConsumer) Close() error {
	return kc.consumer.Close()
}
