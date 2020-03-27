package main

import (
	"github.com/gorilla/mux"
	"log"
	"lungo/pkg"
	"net/http"
)

func main() {
	homeHandler := pkg.Handler{}
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler.Get)
	log.Fatal(http.ListenAndServe(":8080", r))
}