package postgres

import (
	"database/sql"
	"github.com/Blxssy/Todo-app/backend"
	_ "github.com/lib/pq"
	"sync"
)

type Repository struct {
	db *sql.DB
	mu sync.RWMutex
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Get(id int) (backend.Todo, error) {
	return backend.Todo{}, nil
}

func (r *Repository) Put(title string) error {
	return nil
}

func (r *Repository) Latest() ([]backend.Todo, error) {
	return nil, nil
}
