package main

import (
	"github.com/gorilla/mux"
	"log"
	"lungo/handlers"
	"net/http"
	"fmt"
)


func main() {
	httpHandlers := initHandlers()
	r := mux.NewRouter()
	r.HandleFunc("/", httpHandlers.Language.Get)

	fmt.Println("\n - 🚀 Application launched @ http://:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

type Handlers struct {
	Language handlers.LanguageHandler
}

func initHandlers() Handlers {
	return Handlers{
		Language: handlers.LanguageHandler{},
	}
}