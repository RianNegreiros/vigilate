package config

import (
	"crypto/tls"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func NewKafkaWriterConfig() kafka.WriterConfig {
	if os.Getenv("APP_ENV") == "docker" {
		return kafka.WriterConfig{
			Brokers: []string{"kafka:9092"},
			Topic:   "health-check-results",
		}
	}

	topic := os.Getenv("KAFKA_TOPIC")
	broker := os.Getenv("KAFKA_BROKER")
	username := os.Getenv("KAFKA_USERNAME")
	password := os.Getenv("KAFKA_PASSWORD")

	mechanism, err := scram.Mechanism(scram.SHA256, username, password)

	if err != nil {
		log.Fatalln(err)
	}

	dialer := &kafka.Dialer{
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}

	return kafka.WriterConfig{
		Brokers: []string{broker},
		Topic:   topic,
		Dialer:  dialer,
	}
}

func NewKafkaReaderConfig() kafka.ReaderConfig {
	if os.Getenv("APP_ENV") == "docker" {
		return kafka.ReaderConfig{
			Brokers: []string{"kafka:9092"},
			Topic:   "health-check-results",
			GroupID: "health-check-results-consumer",
		}
	}

	topic := os.Getenv("KAFKA_TOPIC")
	broker := os.Getenv("KAFKA_BROKER")
	groupID := os.Getenv("KAFKA_GROUP_ID")
	username := os.Getenv("KAFKA_USERNAME")
	password := os.Getenv("KAFKA_PASSWORD")

	mechanism, err := scram.Mechanism(scram.SHA256, username, password)

	if err != nil {
		log.Fatalln(err)
	}

	dialer := &kafka.Dialer{
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}

	return kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: groupID,
		Dialer:  dialer,
	}
}
