package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// data from user
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// store user in system
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type APIResponse struct {
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
	Success bool   `json:"success"`
}

func writeJSON(response http.ResponseWriter, status int, payload any) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	json.NewEncoder(response).Encode(payload)
}

func writeError(response http.ResponseWriter, status int, message string) {
	writeJSON(response, status, APIResponse{
		Success: false,
		Error:   message,
	})
}

func getHealthhandler(response http.ResponseWriter, request *http.Request) {
	writeJSON(response, http.StatusOK, APIResponse{
		Success: true,
		Data: map[string]string{
			"status":  "ok",
			"version": "1.0.0",
		},
	})
}
func createUserHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writeError(response, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var req CreateUserRequest
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		writeError(response, http.StatusBadRequest, "invaliud JSON body")
		return
	}

	if req.Name == "" || req.Email == "" {
		writeError(response, http.StatusBadRequest, "Fill name and email")
		return
	}

	user := User{
		ID:        1,
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now(),
	}

	writeJSON(response, http.StatusCreated, APIResponse{
		Success: true,
		Data:    user,
	})
}
func getUserHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		writeError(response, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	user := User{
		ID:        1,
		Name:      "Alice",
		Email:     "alice@example.com",
		CreatedAt: time.Now(),
	}
	writeJSON(response, http.StatusOK, APIResponse{
		Success: true,
		Data:    user,
	})
}

func main() {
	http.HandleFunc("/health", getHealthhandler)
	http.HandleFunc("/users", createUserHandler)
	http.HandleFunc("/user", getUserHandler)

	port := ":8080"
	log.Printf("Server starting on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
