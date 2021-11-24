package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Shopify/sarama"
	"github.com/gorilla/mux"
)

var SyncProducer1 sarama.SyncProducer

func PushToKafka(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	msgTopic := mux.Vars(r)["topicname"]

	body, err := io.ReadAll(r.Body)
	stringBody := string(body)

	if err != nil {
		fmt.Fprintf(rw, "something went wrong... could not read body")
		return
	} else {
		fmt.Printf("Body read safely: %s \n", stringBody)
	}

	partitionId, offset, err := SyncProducer1.SendMessage(&sarama.ProducerMessage{
		Topic: msgTopic,
		Value: sarama.StringEncoder(stringBody),
	})

	if err != nil {
		fmt.Printf("error: %v \n", err)
	} else {
		fmt.Fprintf(rw, "message sent, partition id: %v offset: %v \n", partitionId, offset)
		return
	}

}
