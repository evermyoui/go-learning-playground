package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
	//functions
	http.HandleFunc("/health", getHealthHandler)
	http.HandleFunc("/users", createUserHandler)
	http.HandleFunc("/user", getUserHandler)

	//ports
	port := ":8080"
	log.Printf("Server starting on http://localhost%s\n", port)

	// run at port 8080
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

// helpers
func writeJSON(response http.ResponseWriter, status int, payload any) {
	//set the header
	response.Header().Set("Content-Type", "application/json")
	//write status code
	response.WriteHeader(status)
	//write JSON body
	json.NewEncoder(response).Encode(payload)
}

func writeError(response http.ResponseWriter, status int, message string) {
	writeJSON(response, status, APIResponse{
		Success: false,
		Error:   message,
	})
}

// structs
type APIResponse struct {
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
	Success bool   `json:"success"`
}
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Created_At time.Time `json:"created_at"`
}

// handlers
func getHealthHandler(response http.ResponseWriter, request *http.Request) {
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

	//json -> struct
	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		writeError(response, http.StatusBadRequest, "invalid json body")
		return
	}

	if req.Name == "" || req.Email == "" {
		writeError(response, http.StatusBadRequest, "not valid name and email")
		return
	}
	user := User{
		ID:         1,
		Name:       req.Name,
		Email:      req.Email,
		Created_At: time.Now(),
	}

	writeJSON(response, http.StatusOK, APIResponse{
		Success: true,
		Data:    user,
	})
}

func getUserHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		writeError(response, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	//dummy user
	user := User{
		ID:         1,
		Name:       "Alice",
		Email:      "alice@exmaple.com",
		Created_At: time.Now(),
	}

	writeJSON(response, http.StatusOK, APIResponse{
		Success: true,
		Data:    user,
	})
}
