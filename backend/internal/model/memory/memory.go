package memory

import (
	"github.com/Blxssy/Todo-app/backend"
	"github.com/Blxssy/Todo-app/backend/internal/model"
	"sync"
)

type Repository struct {
	data map[int]backend.Todo
	mu   sync.RWMutex
}

func New() *Repository {
	return &Repository{data: make(map[int]backend.Todo)}
}

func (r *Repository) Put(title string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	var todo backend.Todo
	todo.ID = len(r.data)
	todo.Title = title
	todo.State = true

	r.data[todo.ID] = todo

	return nil
}

func (r *Repository) Get(id int) (backend.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var todo backend.Todo

	if _, ok := r.data[id]; !ok {
		return backend.Todo{}, model.ErrNotFound
	}

	return todo, nil
}

func (r *Repository) Latest() ([]backend.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var todos []backend.Todo

	for _, todo := range r.data {
		todos = append(todos, todo)
	}

	return todos, nil
}
