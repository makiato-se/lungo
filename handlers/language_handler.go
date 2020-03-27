package handlers

import (
	"encoding/json"
	"lungo/persistance"
	"net/http"
)

type LanguageHandler struct {
	repository persistance.LanguageRepository
}

func NewLanguageHandler(repository persistance.LanguageRepository) LanguageHandler {
	return LanguageHandler{
		repository: repository,
	}
}

func (handler *LanguageHandler) Get(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	language := query.Get("language")
	wordList := handler.repository.GetWordsForLanguage(language)
	
	words, err := json.Marshal(wordList)
	if err != nil {
		panic(err)
	}
	_, err = writer.Write(words)
	if err != nil {
		panic(err)
	}
	return

	/*writer.WriteHeader(404)
	_, err = writer.Write([]byte("Not Found"))
	if err != nil {
		panic(err)
	}*/
}
