package handlers

import (
	"log"
	"net/http"

	"github.com/dimoynwa/product-api/data"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (p *Products) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		p.getProducts(writer, request)
		return
	} else if request.Method == http.MethodPost {

	}

	writer.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(writer http.ResponseWriter, request *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(writer)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
