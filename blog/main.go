package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bienvenue sur la page d'accueil de mon blog en Go !"))
	}).Methods("GET")

	// Démarrez votre serveur HTTP en écoutant sur le port 8080
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
