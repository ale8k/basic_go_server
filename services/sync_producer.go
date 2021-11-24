package services

import (
	"strings"

	"github.com/Shopify/sarama"
	"github.com/ale8k/basic_go_server/config"
	"github.com/ale8k/basic_go_server/utils"
)

func GetNewSyncProducer() sarama.SyncProducer {
	producer, err := sarama.NewSyncProducer(strings.Split(config.KafkaBrokerAddrs, ","), config.AppSaramaConfig)
	utils.CheckForErrPanic(err)
	return producer
}
