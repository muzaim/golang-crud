package routes

import (
	"golang-crud/controllers/authorcontroller"

	"github.com/gorilla/mux"
)

func AuthorRoutes(r *mux.Router) {
	router := r.PathPrefix("/author").Subrouter()

	router.HandleFunc("", authorcontroller.Index).Methods("GET")
	router.HandleFunc("", authorcontroller.CreateAuthor).Methods("POST")
	router.HandleFunc("/{id}", authorcontroller.GetAuthorByID).Methods("GET")
	router.HandleFunc("/{id}", authorcontroller.UpdateAuthor).Methods("PUT")
	router.HandleFunc("/{id}", authorcontroller.DeleteAuthor).Methods("DELETE")
}
