package src

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	c "github.com/yinkar/tototodo/backend/_config"
)

var config = c.GetConfig()

func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	log.Print("API health is OK.")

	success(w, `{"alive": true}`)
}

func Version(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	success(w, fmt.Sprintf(`{"version": "%s"}`, config.Api.Version))
}

func error(w http.ResponseWriter, content string, status int) {
	w.Header().Set("Content-Type", "application/json")

	data := struct {
		Error string `json:"error"`
	}{Error: content}

	res, _ := json.Marshal(data)
	w.WriteHeader(status)

	fmt.Fprintf(w, string(res))
}

func success(w http.ResponseWriter, content string) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, content)
}
