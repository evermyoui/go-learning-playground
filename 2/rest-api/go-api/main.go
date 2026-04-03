package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	//store and handler
	if err := godotenv.Load(); err != nil {
		log.Println("No env file found")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL env required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	store, err := NewPostgresStore(dsn)
	if err != nil {
		log.Fatalf("could not connect db: %v", err)
	}
	log.Println("Connected to Postgres")

	h := NewUserHandler(store)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/users", func(r chi.Router) { // anonymous function -> func (r chi.Router)
		// r.Get("/", h.ListUsers)
		r.Post("/", h.CreateUser)
		// r.Get("/{id}", h.GetUser)
		// r.Put("/{id}", h.UpdateUser)
		// r.Delete("/{id}", h.DeleteUser)
	})

	log.Printf("Starting server on localhost%s\n", port)

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
