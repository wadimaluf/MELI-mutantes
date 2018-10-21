package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func CheckMutant(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func GetStats(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := mux.NewRouter()

	router.HandleFunc("/mutant", CheckMutant).Methods("POST")
	router.HandleFunc("/stats/", GetStats).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
