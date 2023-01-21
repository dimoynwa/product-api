package data

import (
	"encoding/json"
	"io"
	"time"

	"github.com/dimoynwa/product-api/generator"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"` // Internal Product ID
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(p)
}

type Products []*Product

func (p *Products) ToJSON(writer io.Writer) error {
	return json.NewEncoder(writer).Encode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(prod *Product) {
	id := generator.GenerateProductId()
	prod.ID = id
	productList = append(productList, prod)
}

var productList = Products{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   "",
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   "",
	},
}
