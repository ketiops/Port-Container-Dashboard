package router

import (
	"db-server/controller"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func InitRouter() http.Handler {
	hostIP := os.Getenv("DB_SERVER_HOST")

	r := mux.NewRouter()
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://" + hostIP + ":3000", "http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
	})

	// Container routes
	r.HandleFunc("/container", controller.GetPreContainerInfo).Methods("POST")
	r.HandleFunc("/containers", controller.GetAllPreContainersInfo).Methods("GET")
	r.HandleFunc("/container/spec", controller.CreateContainerSpec).Methods("POST")
	r.HandleFunc("/container/spec", controller.GetAllContainerSpec).Methods("GET")
	r.HandleFunc("/container/spec/review", controller.GetContainerReview).Methods("GET")
	r.HandleFunc("/container/spec/review", controller.AppendContainerReview).Methods("POST")

	//TOS container route
	r.HandleFunc("/container/tos", controller.GetTosContainer).Methods("POST")

	// Security checkpoint routes
	r.HandleFunc("/checkpoint", controller.GetAllCheckpoint).Methods("GET")
	r.HandleFunc("/checkpoint/state", controller.GetAllCheckpointState).Methods("GET")
	r.HandleFunc("/checkpoint/state", controller.CreateCheckpointState).Methods("POST")

	handler := corsConfig.Handler(r)

	return handler
}
