package config

import "github.com/confluentinc/confluent-kafka-go/kafka"

func NewKafkaConfig() *kafka.ConfigMap {
	config :=
		&kafka.ConfigMap{
			"bootstrap.servers": "localhost",
			"group.id":          "healthChecks",
			"auto.offset.reset": "earliest",
		}

	return config
}
