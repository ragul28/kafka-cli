package queue

import (
	"fmt"
	"os"

	"github.com/ragul28/kafka-cli/utils"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Producer(env *utils.Confg) {

	fmt.Println(env)

	config := &kafka.ConfigMap{
		"metadata.broker.list": env.Broker,
		"security.protocol":    env.Protocol,
		"sasl.mechanisms":      env.Mechanisms,
		"sasl.username":        env.Username,
		"sasl.password":        env.Password,
		// "debug":                "generic,broker,security",
	}

	topic := env.Topic

	p, err := kafka.NewProducer(config)
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Producer %v\n", p)

	deliveryChan := make(chan kafka.Event)

	value := "Hello"
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
