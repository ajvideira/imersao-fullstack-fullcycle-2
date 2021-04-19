package main

import (
	"fmt"
	"log"

	proKafka "github.com/ajvideira/imersao-fullstack-fullcycle-2/simulador/application/kafka"
	"github.com/ajvideira/imersao-fullstack-fullcycle-2/simulador/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env: %v", err)
	}
}

func main() {
	/*producer := kafka.NewKafkaProducer();
	kafka.Produce("Olá", "readtest", producer)*/
	fmt.Println("Starting app simulador")

	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()
	
	for msg := range msgChan {
		fmt.Printf("Received message: %v\n", string(msg.Value))
		go proKafka.Produce(msg)
	}

	for {
		_ = 1
	}
}