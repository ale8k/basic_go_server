package config

import (
	"os"

	"github.com/Shopify/sarama"
)

var AppSaramaConfig *sarama.Config
var KafkaBrokerAddrs string

func setSaramaConfig() {
	AppSaramaConfig = sarama.NewConfig()
}

func setKafkaBrokerAddrs() {
	KafkaBrokerAddrs = os.Getenv("KAFKA_BROKER_ADDRS")
}
