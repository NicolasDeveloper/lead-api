package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NicolasDeveloper/lead-api/env"
	"github.com/NicolasDeveloper/lead-api/routes"
	"github.com/gorilla/mux"
)

func main() {
	env.InitConfiguration()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/lead", routes.Post).Methods("POST")
	router.HandleFunc("/lead", routes.Get).Methods("GET")
	router.HandleFunc("/lead/{id}", routes.GetById).Methods("GET")
	router.HandleFunc("/lead/{id}", routes.Put).Methods("PUT")

	fmt.Printf("App running at: %s", env.Parameters.Database.AppPort)
	log.Fatal(http.ListenAndServe(env.Parameters.Database.AppPort, router))

}
