package main

import "net/http"

type UserHandler struct {
	store *UserStore
}

func NewUserHandler(store *UserStore) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

// helpers

// handlers
func (h *UserHandler) ListUsers(response http.ResponseWriter, request *http.Request) {
	users := h.store.List()

	if users == nil {
		users = []User{}
	}
}
