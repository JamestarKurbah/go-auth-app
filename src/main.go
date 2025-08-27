package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	SetupRoutes(r)

	http.Handle("/", r)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	r.HandleFunc("/token", TokenHandler).Methods("GET")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Handler logic for user login
}

func TokenHandler(w http.ResponseWriter, r *http.Request) {
	// Handler logic for token generation
}
