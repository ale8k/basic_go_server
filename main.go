package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Shopify/sarama"

	"github.com/ale8k/basic_go_server/handlers"
	"github.com/ale8k/basic_go_server/routes"
	"github.com/ale8k/basic_go_server/services"
	"github.com/gorilla/mux"
)

var SyncProducer1 sarama.SyncProducer
var Consumer1 sarama.Consumer

func main() {
	router := mux.NewRouter()
	SyncProducer1 = services.GetNewSyncProducer()
	Consumer1 = services.GetNewConsumer()
	handlers.SyncProducer1 = SyncProducer1
	handlers.Consumer1 = Consumer1
	routes.RegisterMessageQueue(router)

	http.Handle("/", router)

	fmt.Println("Server starting on port 8000...")

	log.Fatalf(
		"Server failed to start: %v",
		http.ListenAndServe(":8000", nil),
	)
}
