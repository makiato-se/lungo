package handlers

import "net/http"

type LanguageHandler struct {
	Db int
}

func (handler *LanguageHandler) Get(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	language := query.Get("language")

	if language == "spanish" {
		_, err := writer.Write([]byte("Hola señor!"))
		if err != nil {
			panic(err)
		}
		return
	}

	if language == "french" {
		_, err := writer.Write([]byte("Bonjour monsieur!"))
		if err != nil {
			panic(err)
		}
		return
	}

	writer.WriteHeader(404)
	_, err := writer.Write([]byte("Not Found"))
	if err != nil {
		panic(err)
	}
}