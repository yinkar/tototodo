package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	cors "github.com/rs/cors"
)

var db, _ = gorm.Open("mysql",
	"root:root@/tototodo"+
		"?charset=utf8"+
		"&parseTime=True"+
		"&loc=Local"+
		"&port=33066")

type Todo struct {
	Id      int `gorm:"primary_key"`
	Content string
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todo := &Todo{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&todo)
	if err != nil {
		io.WriteString(w, `{"error": true}`)
	}
	db.Create(&todo)

	result := db.Last(&todo)
	json.NewEncoder(w).Encode(result.Value)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []Todo
	todoRows := db.Find(&todos).Value
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todoRows)
}

func Health(w http.ResponseWriter, r *http.Request) {
	log.Print("API health is OK.")
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func main() {
	defer db.Close()

	db.Debug().DropTableIfExists(&Todo{})
	db.Debug().AutoMigrate(&Todo{})

	log.Print("API serve started.")
	router := mux.NewRouter()

	router.HandleFunc("/health", Health).Methods("GET")
	router.HandleFunc("/todos", CreateTodo).Methods("POST")
	router.HandleFunc("/todos", GetTodos).Methods("GET")

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS"},
	}).Handler(router)

	http.ListenAndServe(":8000", handler)
}
