package queue

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ragul28/kafka-cli/utils"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Producer(cfg *utils.Confg) {

	config := &kafka.ConfigMap{
		"metadata.broker.list": cfg.Broker,
		"security.protocol":    cfg.Protocol,
		"sasl.mechanisms":      cfg.Mechanisms,
		"sasl.username":        cfg.Username,
		"sasl.password":        cfg.Password,
		// "debug":                "generic,broker,security",
	}

	topic := cfg.Topic

	p, err := kafka.NewProducer(config)
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Producer %v\n", p)

	deliveryChan := make(chan kafka.Event)

	msgSize, err := strconv.Atoi(cfg.MsgSize)
	if err != nil {
		log.Fatalln("Error message size not interger", err)
	}

	value := utils.RandString(msgSize)
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(value),
	}, deliveryChan)

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)
}
