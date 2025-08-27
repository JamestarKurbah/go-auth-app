package routes

import (
	"github.com/gorilla/mux"
	"github.com/jamestarkurbah/go-auth-app/src/handlers"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/token", handlers.TokenHandler).Methods("GET")

	return router
}
