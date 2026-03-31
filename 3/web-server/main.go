package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var tasks []Task
var idCounter int

type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"done"`
}

type APIResponse struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, APIResponse{
		Success: false,
		Error:   message,
	})
}

func getTask(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Data:    tasks,
	})
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid json")
		return
	}
	if task.Title == "" {
		writeError(w, http.StatusBadRequest, "Enter Title")
		return
	}

	idCounter++
	newTask := Task{
		ID:     idCounter,
		Title:  task.Title,
		IsDone: false,
	}
	tasks = append(tasks, newTask)
	writeJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Data:    newTask,
	})
}

func updateTask(w http.ResponseWriter, r *http.Request) {

}

func main() {

	r := chi.NewRouter()

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", getTask)
		r.Post("/", createTask)
	})

	port := ":8080"
	log.Printf("Starting server on localhost%s\n", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
