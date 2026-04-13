package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

var BrokerAddress = "localhost:9092"

func SendMessage(topic, key string, value []byte) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{BrokerAddress},
		Topic:   topic,
	})
	defer writer.Close()

	msg := kafka.Message{
		Key:   []byte(key),
		Value: value,
	}
	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Printf("Failed to write message to Kafka: %v", err)
		return err
	}
	return nil
}
