package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Blxssy/Todo-app/backend"
	"github.com/Blxssy/Todo-app/backend/internal/controller"
	"github.com/gorilla/mux"
)

type Handler struct {
	ctrl *controller.Controller
}

func New(ctrl *controller.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	todos, err := h.ctrl.Repo.Latest()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(todos)
	// for _, t := range todos {
	// 	json.NewEncoder(w).Encode(t)
	// }
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

func (h *Handler) CompleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}
	todos, err := h.ctrl.Repo.Latest()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].State = true
			json.NewEncoder(w).Encode(todos[i])
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}

func (h *Handler) HandleDeleteTodo(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing todo ID", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	err = h.ctrl.Repo.Delete(id)
	if err != nil {
		http.Error(w, "Failed to delete todo", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
