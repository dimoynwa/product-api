package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/dimoynwa/product-api/data"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (p *Products) GetProducts(writer http.ResponseWriter, request *http.Request) {
	p.logger.Println("Handle GET products request")
	lp := data.GetProducts()
	err := lp.ToJSON(writer)
	if err != nil {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (p *Products) AddProduct(writer http.ResponseWriter, request *http.Request) {
	p.logger.Println("Handle POST products request")

	prod := request.Context().Value(KeyProduct{}).(*data.Product)

	p.logger.Printf("Going to store: %v\n", prod)
	data.AddProduct(prod)
	writer.WriteHeader(http.StatusCreated)
}

func (p *Products) UpdateProduct(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(writer, "Unable to convert ID", http.StatusBadRequest)
		return
	}

	p.logger.Println("Handle PUT products request with ID ", id)

	prod := request.Context().Value(KeyProduct{}).(*data.Product)

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(writer, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(writer, "Unexpected error", http.StatusInternalServerError)
	}
}

type KeyProduct struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		prod := &data.Product{}

		err := prod.FromJSON(request.Body)
		if err != nil {
			http.Error(writer, "cannot deseriaze Product object", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(request.Context(), KeyProduct{}, prod)
		request = request.WithContext(ctx)
		next.ServeHTTP(writer, request)
	})
}
