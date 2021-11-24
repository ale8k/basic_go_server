package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Shopify/sarama"

	"github.com/ale8k/basic_go_server/config"

	"github.com/ale8k/basic_go_server/routes"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	routes.RegisterPushToKafkaRoutes(router)
	producer, _ := sarama.NewSyncProducer([]string{config.KafkaBrokerAddrs}, sarama.NewConfig())

	defer producer.Close()
	// router.
	// 	Path("/testing").
	// 	Methods("GET").
	// 	HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	// 		rw.Write([]byte("test"))
	// 	})

	// router.HandleFunc("/hi/{pathvar:[a-z]}/dude", func(rw http.ResponseWriter, r *http.Request) {
	// 	pathVarMap := mux.Vars(r)

	// 	fmt.Println("yo", pathVarMap["pathvar"], r.URL.Query().Get("yolo"))
	// 	rw.Write([]byte("hi"))
	// })

	http.Handle("/", router)

	fmt.Println("Server starting on port 8000...")
	// Does it default to hostname of container?
	log.Fatalf(
		"Server failed to start: %v",
		http.ListenAndServe(":8000", nil),
	)
}
