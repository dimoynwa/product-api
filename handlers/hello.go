package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	logger *log.Logger
}

func NewHello(logger *log.Logger) *Hello {
	return &Hello{logger: logger}
}

func (h *Hello) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	h.logger.Println("Hello, Dimo")
	data, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Ooopaaa", http.StatusBadRequest)
		return
	}
	h.logger.Printf("Data readed : %s\n", data)

	fmt.Fprintf(writer, "Hello from server\n")

}
