package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	store UserStorer
}

func NewUserHandler(store UserStorer) *UserHandler {
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

func parseID(request *http.Request) (int, error) {
	return strconv.Atoi(chi.URLParam(request, "id"))
}

// handlers
func (h *UserHandler) ListUsers(response http.ResponseWriter, request *http.Request) {
	users, err := h.store.List()
	if err != nil {
		log.Printf("Error: %v ", err)
		writeError(response, http.StatusInternalServerError, "failed to list users")
		return
	}

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

	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		writeError(response, http.StatusBadRequest, "invalid JSON")
		return
	}

	if req.Email == "" || req.Name == "" {
		writeError(response, http.StatusBadRequest, "fill all fields")
		return
	}

	user, err := h.store.Create(req.Name, req.Email)
	if err != nil {
		writeError(response, http.StatusInternalServerError, "failed to create user")
		return
	}
	writeJSON(response, http.StatusCreated, APIResponse{
		Success: true,
		Data:    user,
	})
}

func (h *UserHandler) GetUser(response http.ResponseWriter, request *http.Request) {
	id, err := parseID(request)
	if err != nil {
		writeError(response, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := h.store.GetByID(id)
	if err != nil {
		writeError(response, http.StatusNotFound, err.Error())
		return
	}

	writeJSON(response, http.StatusOK, APIResponse{
		Success: true,
		Data:    user,
	})
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	var req UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON Body")
		return
	}

	name, email := "", ""

	if req.Name != nil {
		name = *req.Name
	}

	if req.Email != nil {
		email = *req.Email
	}

	user, err := h.store.Update(id, name, email)

	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Data:    user,
	})
}

// func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
// 	id, err := parseID(r)
// 	if err != nil {
// 		writeError(w, http.StatusBadRequest, "invalid id")
// 		return
// 	}

// 	if err := h.store.Delete(id); err != nil {
// 		writeError(w, http.StatusNotFound, err.Error())
// 		return
// 	}

// 	writeJSON(w, http.StatusOK, APIResponse{
// 		Success: true,
// 		Data:    "User Deleted",
// 	})
// }
