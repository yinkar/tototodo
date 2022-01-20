package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	c "github.com/yinkar/tototodo/backend/_config"
	"github.com/yinkar/tototodo/backend/src"
)

var port int

func main() {
	config := c.GetConfig()
	port = config.Srv.Port

	router := mux.NewRouter()

	router.HandleFunc("/version", src.Version).Methods("GET")
	router.HandleFunc("/health", src.Health).Methods("GET")

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
	}).Handler(router)

	log.Printf("API is serving on %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
	if err != nil {
		log.Fatal(err)
	}
}
