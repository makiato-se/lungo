package main

import (
	"github.com/gorilla/mux"
	"log"
	"lungo/pkg"
	"net/http"
)

type Handlers struct {
	Home pkg.HomeHandler
}

func main() {
	handlers := initHandlers()
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Home.Get)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func initHandlers() Handlers {
	return Handlers{
		Home: pkg.HomeHandler{},
	}
}