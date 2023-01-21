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
		p.addProduct(writer, request)
		return
	}

	writer.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(writer http.ResponseWriter, request *http.Request) {
	p.logger.Println("Handle GET products request")
	lp := data.GetProducts()
	err := lp.ToJSON(writer)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (p *Products) addProduct(writer http.ResponseWriter, request *http.Request) {
	p.logger.Println("Handle POST products request")

	prod := &data.Product{}

	err := prod.FromJSON(request.Body)
	if err != nil {
		http.Error(writer, "cannot deseriaze Product object", http.StatusBadRequest)
		return
	}

	p.logger.Printf("Going to store: %v\n", prod)
	data.AddProduct(prod)
	writer.WriteHeader(http.StatusCreated)
}
