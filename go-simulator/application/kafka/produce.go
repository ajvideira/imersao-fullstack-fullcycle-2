package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/ajvideira/imersao-fullstack-fullcycle-2/simulador/application/route"
	"github.com/ajvideira/imersao-fullstack-fullcycle-2/simulador/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)


func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()

	log.Println(string(msg.Value))

	route := &route.Route{}

	json.Unmarshal(msg.Value, route)

	log.Println(route.ClientID)

	route.LoadPositions()

	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Fatalf("Error exporting positions: %v", err)
	}

	for _, pos := range positions {
		kafka.Produce(pos, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}