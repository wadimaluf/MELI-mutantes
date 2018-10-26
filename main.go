package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type MutantRequestBody struct {
	Dna [6]string
}

func CheckMutant(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var b MutantRequestBody

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if isMutant(b.Dna) {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusForbidden)
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

	router.HandleFunc("/mutant/", CheckMutant).Methods("POST")
	router.HandleFunc("/stats/", GetStats).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
