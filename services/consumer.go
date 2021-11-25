package services

import (
	"strings"

	"github.com/Shopify/sarama"
	"github.com/ale8k/basic_go_server/config"
	"github.com/ale8k/basic_go_server/utils"
)

func GetNewConsumer() sarama.Consumer {
	consumer, err := sarama.NewConsumer(strings.Split(config.KafkaBrokerAddrs, ","), config.AppSaramaConfig)
	utils.CheckForErrPanic(err)
	return consumer
}
