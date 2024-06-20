package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Blxssy/Todo-app/backend"
	"github.com/Blxssy/Todo-app/backend/internal/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type Repository struct {
	db *sqlx.DB
	mu sync.RWMutex
}

func New(cfg Config) (*Repository, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}

func (r *Repository) CreateTables() error {
	_, err := r.db.Exec(`CREATE TABLE IF NOT EXISTS Todos (
        id SERIAL PRIMARY KEY,
        title text NOT NULL,
        status boolean NOT NULL
    );`)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *Repository) Close() {
	r.db.Close()
}

func (r *Repository) Get(id int) (backend.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var todo backend.Todo
	query := `SELECT title, status FROM todos WHERE id = $1;`
	err := r.db.QueryRow(query, id).Scan(&todo.Title, &todo.State)
	if err != nil {
		if err == sql.ErrNoRows {
			return todo, model.ErrNotFound
		}
		return todo, err
	}

	return todo, nil
}

func (r *Repository) Put(title string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	query := `INSERT INTO todos (title, status) VALUES ($1, $2);`
	_, err := r.db.Exec(query, title, false)
	if err != nil {
		return fmt.Errorf("could not insert todo: %v", err)
	}

	return nil
}

func (r *Repository) Latest() ([]backend.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	err := r.db.Ping()
	if err != nil {
		return nil, err
	}

	query := `SELECT id, title, status FROM todos ORDER BY id DESC LIMIT 10;`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("could not fetch latest todos: %v", err)
	}
	defer rows.Close()

	var todos []backend.Todo
	for rows.Next() {
		var todo backend.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.State); err != nil {
			return nil, fmt.Errorf("could not scan row: %v", err)
		}
		if !todo.State {
			todos = append(todos, todo)
		}

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *Repository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	log.Println(id)
	query := `DELETE FROM todos WHERE id = $1;`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
