package routes

import (
	"github.com/ale8k/basic_go_server/handlers"
	"github.com/gorilla/mux"
)

func RegisterPushToKafkaRoutes(r *mux.Router) {
	router := r.PathPrefix("/kafka").Subrouter()
	router.HandleFunc("/push/{topicname}", handlers.PushToKafka)
}
