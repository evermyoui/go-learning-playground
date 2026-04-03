package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
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
