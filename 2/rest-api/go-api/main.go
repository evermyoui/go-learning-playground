package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	//store and handler
	store := NewUserStore()
	h := NewUserHandler(store) // dependency injection

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/users", func(r chi.Router) { // anonymous function -> func (r chi.Router)
		r.Get("/", h.ListUsers) //h.ListUsers correct because you're passing the function, not calling it.
		r.Post("/", h.CreateUser)
		r.Get("/{id}", h.GetUser)
		r.Put("/{id}", h.UpdateUser)
		r.Delete("/{id}", h.DeleteUser)
	})

	port := ":8080"
	log.Printf("Starting server on localhost%s\n", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
