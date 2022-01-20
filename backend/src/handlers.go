package src

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	gorm "github.com/jinzhu/gorm"
	c "github.com/yinkar/tototodo/backend/_config"
)

var config = c.GetConfig()

type ApiHandlers struct {
	Db *gorm.DB
}

type Todo struct {
	Id      int `gorm:"primary_key"`
	Content string
}

// GET /health
func (a *ApiHandlers) Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		a.error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	log.Print("API health is OK.")

	a.success(w, `{"alive": true}`)
}

// GET /version
func (a *ApiHandlers) Version(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		a.error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	a.success(w, fmt.Sprintf(`{"version": "%s"}`, config.Api.Version))
}

// GET /todos
func (a *ApiHandlers) GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []Todo
	todoRows := a.Db.Find(&todos).Value
	output, _ := json.Marshal(todoRows)
	a.success(w, string(output))
}

// POST /todos
func (a *ApiHandlers) CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todo := &Todo{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&todo)
	if err != nil {
		io.WriteString(w, `{"error": true}`)
	}
	a.Db.Create(&todo)

	result := a.Db.Last(&todo)

	output, _ := json.Marshal(result.Value)
	a.success(w, string(output))
}

func (a *ApiHandlers) error(w http.ResponseWriter, content string, status int) {
	w.Header().Set("Content-Type", "application/json")

	data := struct {
		Error string `json:"error"`
	}{Error: content}

	res, _ := json.Marshal(data)
	w.WriteHeader(status)

	fmt.Fprintf(w, string(res))
}

func (a *ApiHandlers) success(w http.ResponseWriter, content string) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, content)
}
