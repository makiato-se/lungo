package pkg

import "net/http"

type HomeHandler struct {
	Db int
}

func (handler *HomeHandler) Get(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("Hello World!"))
	if err != nil {
		panic(err)
	}
}