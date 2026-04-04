package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// connect to postgres db
type PostgresStore struct {
	db *sqlx.DB
}

func NewPostgresStore(dsn string) (*PostgresStore, error) {
	//open connection postgres -> sqlx
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect %w", err)
	}

	db.SetMaxOpenConns(25) // max 25 connection to db
	db.SetMaxIdleConns(5)  // max idle connection

	return &PostgresStore{db: db}, nil
}

// create function
func (s *PostgresStore) Create(name, email string) (User, error) {
	var user User

	query := `
			INSERT INTO users (name, email)
			VALUES ($1, $2)
			RETURNING *
	`
	if err := s.db.Get(&user, query, name, email); err != nil {
		return User{}, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}

func (s *PostgresStore) List() ([]User, error) {
	var users []User

	query := `
			SELECT *
			FROM users
	`
	if err := s.db.Select(&users, query); err != nil {
		return nil, fmt.Errorf("list users: %w", err)
	}

	if users == nil {
		users = []User{}
	}

	return users, nil
}
