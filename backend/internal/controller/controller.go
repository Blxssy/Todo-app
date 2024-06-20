package controller

import "github.com/Blxssy/Todo-app/backend"

type todoRepository interface {
	Get(id int) (backend.Todo, error)
	Put(title string) error
	Latest() ([]backend.Todo, error)
	Delete(id int) error
}

type Controller struct {
	Repo todoRepository
}

func New(repo todoRepository) *Controller {
	return &Controller{Repo: repo}
}
