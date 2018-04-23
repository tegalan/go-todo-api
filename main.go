package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Welcome to ToDo App")

}

func TodoIndex(w http.ResponseWriter, r *http.Request) {

	todos := Todos{
		Todo{Name: "Write Presentation"},
		Todo{Name: "Host Meetup"},
	}

	json.NewEncoder(w).Encode(todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoId"]

	fmt.Fprintln(w, "Todo Show", todoID)
}
