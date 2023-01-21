package handlers

import (
	"log"
	"net/http"

	"github.com/dimoynwa/product-api/data"
	"github.com/dimoynwa/product-api/helpers"
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
	} else if request.Method == http.MethodPut {
		id, err := helpers.ExtractIdFromPath(request.URL.Path)
		if err != nil {
			p.logger.Printf("Error extracting ID from request: %v\n", err)
			http.Error(writer, "Not valid URL", http.StatusBadRequest)
			return
		}
		p.logger.Printf("Extracted Product ID : %v\n", id)
		p.updateProduct(id, writer, request)
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

func (p *Products) updateProduct(id int, writer http.ResponseWriter, request *http.Request) {
	p.logger.Println("Handle PUT products request with ID ", id)

	prod := &data.Product{}

	err := prod.FromJSON(request.Body)
	if err != nil {
		http.Error(writer, "cannot deseriaze Product object", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(writer, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(writer, "Unexpected error", http.StatusInternalServerError)
	}
}
