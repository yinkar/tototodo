package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/yinkar/tototodo/backend/src"
)

var port int = 8000

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/version", src.Version).Methods("GET")
	router.HandleFunc("/health", src.Health).Methods("GET")

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
	}).Handler(router)

	log.Printf("API is serving on %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
}
