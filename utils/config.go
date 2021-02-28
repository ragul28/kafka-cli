package utils

import (
	"flag"
	"os"
)

type Confg struct {
	Broker     string
	Protocol   string
	Mechanisms string
	Username   string
	Password   string
	Groupid    string
	Topic      string
	MsgSize    string
}

func LoadConfig() *Confg {
	var c Confg

	c.Broker = getConfig("KAFKA_BROKERS", "broker", "localhost:9200", "Kafka broker url")
	c.Protocol = getConfig("KAFKA_PROTOCOL", "protocol", "PLAINTEXT", "Kafka broker protocol")
	c.Mechanisms = getConfig("KAFKA_SASLMEC", "SASL", "PLAIN", "Kafka SASL")
	c.Username = getConfig("KAFKA_USERNAME", "user", "", "Kafka broker user")
	c.Password = getConfig("KAFKA_PASSWORD", "password", "", "Kafka broker password")
	c.Groupid = getConfig("KAFKA_GROUPID", "gid", "gid", "Kafka group id")
	c.Topic = getConfig("KAFKA_TOPIC", "topic", "topic", "Kafka topic name")
	c.MsgSize = getConfig("KAFKA_MSGSIZE", "msgsize", "10", "Kafka message size")

	return &c
}

func getConfig(key, flagStr, defaultVal, flagUsage string) string {

	flagValue := flag.String(flagStr, "", flagUsage)

	if len(*flagValue) == 0 {
		sysEnv := os.Getenv(key)
		if len(sysEnv) == 0 {
			return defaultVal
		}
		return sysEnv
	}

	return *flagValue
}
