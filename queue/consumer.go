package queue

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ragul28/kafka-cli/utils"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Consumer(cfg *utils.Confg) {

	config := &kafka.ConfigMap{
		"metadata.broker.list":            cfg.Broker,
		"security.protocol":               cfg.Protocol,
		"sasl.mechanisms":                 cfg.Mechanisms,
		"sasl.username":                   cfg.Username,
		"sasl.password":                   cfg.Password,
		"group.id":                        cfg.Groupid,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"default.topic.config":            kafka.ConfigMap{"auto.offset.reset": "earliest"},
		//"debug":                           "generic,broker,security",
	}

	topic := cfg.Topic

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	c, err := kafka.NewConsumer(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Created Consumer %v\n", c)
	err = c.Subscribe(topic, nil)
	run := true
	counter := 0
	commitAfter := 1000
	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		case ev := <-c.Events():
			switch e := ev.(type) {
			case kafka.AssignedPartitions:
				c.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				c.Unassign()
			case *kafka.Message:
				fmt.Printf("%% Message on %s: %s\n", e.TopicPartition, string(e.Value))
				counter++
				if counter > commitAfter {
					c.Commit()
					counter = 0
				}

			case kafka.PartitionEOF:
				fmt.Printf("%% Reached %v\n", e)
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
				run = false
			}
		}
	}
	fmt.Printf("Closing consumer\n")
	c.Close()
}
