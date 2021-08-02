package main

import (
	_ "encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	r.Handle("/", http.FileServer(http.Dir("./views/")))

	r.Handle("/status", NotImplemented).Methods("GET")
	r.Handle("/products", NotImplemented).Methods("GET")
	r.Handle("/products/{slug}/feedback", NotImplemented).Methods("POST")

	http.ListenAndServe(":8000", r)

}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})
