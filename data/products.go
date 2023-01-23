package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/dimoynwa/product-api/generator"
	"github.com/dimoynwa/product-api/validation"
	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"` // Internal Product ID
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

func (p *Product) Validate() error {
	validator := validator.New()
	validator.RegisterValidation("sku", validation.ValidateSku)

	err := validator.Struct(p)
	return err
}

func GetProducts() Products {
	return productList
}

func AddProduct(prod *Product) {
	id := generator.GenerateProductId()
	prod.ID = id
	productList = append(productList, prod)
}

func UpdateProduct(id int, prod *Product) error {
	fp, i, err := findProduct(id)
	if err != nil {
		return err
	}

	prod.ID = fp.ID
	productList[i] = prod
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
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
