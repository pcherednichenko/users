package handler

import (
	"net/http"

	_ "github.com/pcherednichenko/users/docs"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

// Router contains all possible routes of our service to handle
func (s *server) Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/user", s.createUser).Methods("POST")

	router.HandleFunc("/user/{id:[0-9]+}", s.getUser).Methods("GET")
	router.HandleFunc("/user/{id:[0-9]+}", s.updateUser).Methods("PUT")
	router.HandleFunc("/user/{id:[0-9]+}", s.deleteUser).Methods("DELETE")

	router.HandleFunc("/users", s.getUsers).Methods("GET")

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	return router
}
