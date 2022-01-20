package src

import (
	"io"
	"log"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	log.Print("API health is OK.")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func Version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"version": "v0.1.0"}`)
}
