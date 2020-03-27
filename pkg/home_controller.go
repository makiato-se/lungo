package pkg

import "net/http"

type Handler struct {

}

func (handler *Handler) Get(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("Hello World!"))
	if err != nil {
		panic(err)
	}
}