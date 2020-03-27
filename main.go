package main

import (
	"github.com/gorilla/mux"
	"log"
	"lungo/pkg"
	"net/http"
	"fmt"
)

func main() {
	homeHandler := pkg.Handler{}
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler.Get)

	fmt.Println("\n - 🚀 Application launched @ http://:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}