package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	route "github.com/raul-marques/imersaofsfc2-simulator/application/route"
)

func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
	}
	p, err := ckafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

func Publish(msg string, topic string, producer *ckafka.Producer) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartitions{Topic: &topic, Partition: ckafka.PartitionAny},
	}
}

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route2 := route.NewRoute()
	json.Unmarshal(msg.Value, &route2)
	route2.LoadPositions()
	positions, err := route2.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
