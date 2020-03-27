package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	fmt.Println("\n - 🚀 Application launched @ http://:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func HomeHandler(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("Hello World!"))
	if err != nil {
		panic(err)
	}
}
