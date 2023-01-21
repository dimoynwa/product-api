package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	logger *log.Logger
}

func NewGoodbye(logger *log.Logger) *Goodbye {
	return &Goodbye{logger}
}

func (g *Goodbye) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Byeee\n"))
}
