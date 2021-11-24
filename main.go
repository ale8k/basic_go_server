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

func main() {
	router := mux.NewRouter()
	SyncProducer1 = services.GetNewSyncProducer()
	handlers.SyncProducer1 = SyncProducer1
	routes.RegisterMessageQueue(router)

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
