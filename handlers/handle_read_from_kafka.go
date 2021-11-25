package handlers

import (
	"net/http"
	"strconv"

	"github.com/Shopify/sarama"
	"github.com/gorilla/mux"
)

var Consumer1 sarama.Consumer

func ReadFromKafka(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	handleGenericErr := func(err error) {
		if err != nil {
			rw.Write([]byte("Something went wrong..."))
		}
	}
	msgTopic := mux.Vars(r)["topicname"]
	offset, err := strconv.Atoi(r.URL.Query().Get("lowerBounds"))
	handleGenericErr(err)

	readCount, err := strconv.Atoi(r.URL.Query().Get("higherBounds"))
	handleGenericErr(err)

	consumer, err := Consumer1.ConsumePartition(msgTopic, 0, int64(offset))
	handleGenericErr(err)

	respStr := ""

	for readCount != 0 {
		readCount--
		msg := <-consumer.Messages()
		respStr += string(msg.Value)
	}

	rw.Write([]byte(respStr))

}
