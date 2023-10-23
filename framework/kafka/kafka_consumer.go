package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	consumer *kafka.Consumer
}

func NewKafkaConsumer(consumer *kafka.Consumer) *KafkaConsumer {
	return &KafkaConsumer{consumer}
}

func (kp *KafkaConsumer) ConsumeMessages(topic string, messageHandler func([]byte) error) error {
	kp.consumer.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := kp.consumer.ReadMessage(-1)
		if err == nil {
			if err := messageHandler(msg.Value); err != nil {
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
