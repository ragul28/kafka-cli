package main

import (
	"github.com/ragul28/kafka-cli/queue"
	"github.com/ragul28/kafka-cli/utils"
)

func main() {

	cfg := utils.LoadConfig()

	queue.Producer(cfg)
	queue.Consumer(cfg)
}
