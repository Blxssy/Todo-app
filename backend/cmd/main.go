package main

import (
	"net/http"

	"github.com/Blxssy/Todo-app/backend/internal/controller"
	httphandler "github.com/Blxssy/Todo-app/backend/internal/handler/http"
	"github.com/Blxssy/Todo-app/backend/internal/model/memory"
	"github.com/gorilla/mux"
)

func main() {
	staticDir := "./frontend/static"

	router := mux.NewRouter()

	repo := memory.New()
	ctrl := controller.New(repo)
	h := httphandler.New(ctrl)

	router.HandleFunc("/todos", h.Home).Methods("GET")
	router.HandleFunc("/todos", h.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}/complete", h.CompleteTodo).Methods("PUT")
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(staticDir))))

	http.ListenAndServe(":8080", router)
}
