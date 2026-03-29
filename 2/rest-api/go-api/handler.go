package main

import (
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	store *UserStore
}

func NewUserHandler(store *UserStore) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

// helpers

func writeJSON(response http.ResponseWriter, status int, payload any) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	json.NewEncoder(response).Encode(payload)
}

func writeError(response http.ResponseWriter, status int, message string) {
	writeJSON(response, status, APIResponse{
		Error:   message,
		Success: false,
	})
}

// handlers
func (h *UserHandler) ListUsers(response http.ResponseWriter, request *http.Request) {
	users := h.store.List()

	if users == nil {
		users = []User{}
	}
	writeJSON(response, http.StatusOK, APIResponse{
		Success: true,
		Data:    users,
	})
}

func (h *UserHandler) CreateUser(response http.ResponseWriter, request *http.Request) {
	var req CreateUserRequest

	if err := json.NewDecoder(request.Body).Decode(&request); err != nil {
		writeError(response, http.StatusBadRequest, "invalid JSON")
		return
	}

	if req.Email == "" || req.Name == "" {
		writeError(response, http.StatusBadRequest, "fill all fields")
	}

	user := h.store.Create(req.Name, req.Email)
	writeJSON(response, http.StatusCreated, APIResponse{
		Success: true,
		Data:    user,
	})
}
