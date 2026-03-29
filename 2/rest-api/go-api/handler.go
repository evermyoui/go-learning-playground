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

func errorJSON(response http.ResponseWriter, status int, message string) {
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
}
