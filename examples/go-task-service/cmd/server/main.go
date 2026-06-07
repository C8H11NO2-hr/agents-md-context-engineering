package main

import (
	"encoding/json"
	"log"
	"net/http"

	task "github.com/example/agents-md-context-engineering/examples/go-task-service"
)

func main() {
	store := task.NewStore()
	store.Create("add an AGENTS.md file close to the code")

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(store.List()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	log.Println("listening on http://localhost:8080/tasks")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
