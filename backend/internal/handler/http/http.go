package http

import (
	"encoding/json"
	"github.com/Blxssy/Todo-app/backend"
	"github.com/Blxssy/Todo-app/backend/internal/controller"
	"net/http"
)

type Handler struct {
	ctrl *controller.Controller
}

func New(ctrl *controller.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	todos, err := h.ctrl.Repo.Latest()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	for _, t := range todos {
		json.NewEncoder(w).Encode(t)
	}
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t backend.Todo
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	defer r.Body.Close()
	h.ctrl.Repo.Put(t.Title)
}
