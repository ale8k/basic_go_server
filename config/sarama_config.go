package config

import (
	"os"

	"github.com/Shopify/sarama"
)

func setSaramaConfig() {
	AppSaramaConfig = sarama.NewConfig()
	AppSaramaConfig.Producer.Return.Successes = true
}

func setKafkaBrokerAddrs() {
	KafkaBrokerAddrs = os.Getenv("KAFKA_BROKER_ADDRS")
}
