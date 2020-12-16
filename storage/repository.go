package storage

import "database/sql"

// Repository structure to c
type Repository struct {
	db *sql.DB
}

// New returns a new repository with relevant methods configured
func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// DoSomething function is a method of Repository
func (r *Repository) DoSomething() error {
	return nil
}

// use sqlc for generating db access code in Go
