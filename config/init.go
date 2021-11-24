package config

import (
	"fmt"
	"os"

	"github.com/Shopify/sarama"
)

var AppSaramaConfig *sarama.Config
var KafkaBrokerAddrs string

func init() {
	fmt.Fprint(os.Stdin, "initialising sarama config \n")
	setSaramaConfig()
	fmt.Fprint(os.Stdin, "setting kafka address to global var \n")
	setKafkaBrokerAddrs()
}
