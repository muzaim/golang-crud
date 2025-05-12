package routes

import (
	"golang-crud/controllers/bookcontroller"

	"github.com/gorilla/mux"
)

func BookRoutes(r *mux.Router) {
	router := r.PathPrefix("/book").Subrouter()

	router.HandleFunc("", bookcontroller.Index).Methods("GET")
	router.HandleFunc("/{id}", bookcontroller.GetBookByID).Methods("GET")
	router.HandleFunc("/create", bookcontroller.Create).Methods("POST")
	router.HandleFunc("/update/{id}", bookcontroller.Update).Methods("PUT")
	router.HandleFunc("/delete/{id}", bookcontroller.Delete).Methods("DELETE")
}
