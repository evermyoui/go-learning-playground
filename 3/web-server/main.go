package main

import (
	"encoding/json"
	"fmt"
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

type CreateTaskRequest struct {
	Title string `json:"title"`
}

type UpdateTaskRequest struct {
	Title  *string `json:"title"`
	IsDone *bool   `json:"done"`
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

func getTasks(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Data:    tasks,
	})
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var task CreateTaskRequest
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
	writeJSON(w, http.StatusCreated, APIResponse{
		Success: true,
		Data:    newTask,
	})
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	var input UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	for i, task := range tasks {
		if task.ID == id {
			if input.Title == nil && input.IsDone == nil {
				writeError(w, http.StatusBadRequest, "Nothing update")
				return
			}
			if input.Title != nil {
				tasks[i].Title = *input.Title
			}
			if input.IsDone != nil {
				tasks[i].IsDone = *input.IsDone
			}
			writeJSON(w, http.StatusOK, APIResponse{
				Success: true,
				Data:    tasks[i],
			})
		}
	}
	writeError(w, http.StatusBadRequest, "Task not found")
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			writeJSON(w, http.StatusOK, APIResponse{
				Success: true,
				Data:    "Task Deleted",
			})
			return
		}
	}
	writeError(w, http.StatusNotFound, "No ID found")
}

func getTask(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		writeError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	for _, task := range tasks {
		if task.ID == id {
			writeJSON(w, http.StatusOK, APIResponse{
				Success: true,
				Data:    task,
			})
			return
		}
	}
	writeError(w, http.StatusNotFound, "No ID found")
}

func main() {

	r := chi.NewRouter()

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", getTasks)
		r.Get("/{id}", getTask)
		r.Post("/", createTask)
		r.Put("/{id}", updateTask)
		r.Delete("/{id}", deleteTask)
	})

	port := ":8080"
	log.Printf("Starting server on localhost%s\n", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
