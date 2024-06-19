package main

import (
	"github.com/Blxssy/Todo-app/backend/internal/controller"
	httphandler "github.com/Blxssy/Todo-app/backend/internal/handler/http"
	"github.com/Blxssy/Todo-app/backend/internal/model/memory"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	//router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",))
	//router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	repo := memory.New()
	ctrl := controller.New(repo)
	h := httphandler.New(ctrl)

	router.HandleFunc("/", h.Home).Methods("GET")
	router.HandleFunc("/todo", h.CreateTodo).Methods("POST")

	http.ListenAndServe(":8080", router)
}
