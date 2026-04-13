package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func StartBookingConsumer(BrokerAddress, topic,	groupId string ) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{BrokerAddress},
		GroupID: groupId,
		Topic:   topic,
		MinBytes: 10e3,  // 10KB
		MaxBytes: 10e6,  // 10MB
	})
	log.Printf("[Kafka Consumer] Listening to topic '%s'...", topic)
	
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}
		fmt.Printf("\n📩 Received message from topic [%s]:\nKey: %s\nValue: %s\n", m.Topic, string(m.Key), string(m.Value))

	}

}