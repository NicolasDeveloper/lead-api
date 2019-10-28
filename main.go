package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NicolasDeveloper/go-rest-api/routes"
	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/lead", routes.Post).Methods("POST")
	router.HandleFunc("/lead", routes.Get).Methods("GET")
	router.HandleFunc("/lead/{id}", routes.GetById).Methods("GET")
	router.HandleFunc("/lead/{id}", routes.Put).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3002", router))
}
