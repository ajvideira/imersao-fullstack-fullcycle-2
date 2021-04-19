package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	MsgChan chan *ckafka.Message
}

func NewKafkaConsumer(msgChannel chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MsgChan: msgChannel,
	}
}

func (k *KafkaConsumer) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id": os.Getenv("KafkaConsumerGroupId"),
		"security.protocol": os.Getenv("security.protocol"),
		"sasl.mechanisms": os.Getenv("sasl.mechanisms"),
		"sasl.username": os.Getenv("sasl.username"),
		"sasl.password": os.Getenv("sasl.password"),
	}

	consumer, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("Can't create kafka consumer: %v", err)
	}

	topics := []string{os.Getenv("KafkaReadTopic")}

	consumer.SubscribeTopics(topics, nil)
	fmt.Println("Kafka Consumer has been started")

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Fatalf("Error receiving message: %v", err)
		} else {
			k.MsgChan <- msg
		}
	}
}