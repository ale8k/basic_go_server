package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func PushToKafka(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("pushing to kafka"))

	msgTopic := mux.Vars(r)["topicname"]

	fmt.Printf("Topic msg sent to: %s \n", msgTopic)

	// producer, _ := sarama.NewSyncProducer([]string{config.KafkaBrokerAddrs}, sarama.NewConfig())
	// producer.SendMessage(&sarama.ProducerMessage{Topic: "yololo", Value: sarama.StringEncoder("hi")})
	// producer, _ := sarama.NewAsyncProducer([]string{config.KafkaBrokerAddrs}, sarama.NewConfig())
	// producer.Input() <- &sarama.ProducerMessage{Topic: "gololo", Value: sarama.StringEncoder("boop")}

}
