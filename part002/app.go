package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", hello)
	r.HandleFunc("/readiness", checkReadiness).Methods("GET")
	r.HandleFunc("/liveness", checkLiveness).Methods("GET")

	http.ListenAndServe(":3000", r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello world!"))
}

func checkReadiness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("I am ready!"))
}

func checkLiveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("I am alive!"))
}
