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

	db := src.Connect()

	apiHandlers := &src.ApiHandlers{Db: db}

	router.HandleFunc("/version", apiHandlers.Version).Methods("GET")
	router.HandleFunc("/health", apiHandlers.Health).Methods("GET")
	router.HandleFunc("/todos", apiHandlers.GetTodos).Methods("GET")
	router.HandleFunc("/todos", apiHandlers.CreateTodo).Methods("POST")

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
	}).Handler(router)

	// Listen the port
	log.Printf("API is serving on %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
	if err != nil {
		log.Fatal(err)
	}
}
