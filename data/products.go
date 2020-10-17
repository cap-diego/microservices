package data

import (
	"fmt"
)

// Product defines the structure for an API
// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// the name for this poduct
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"description"`

	// the price for the product
	//
	// required: true
	// min: 0.01
	Price float64 `json:"price" validate:"required,gt=0"`

	// the SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"sku"`
}

//Products is a collection a product
type Products []*Product

// GetProducts returns a list of products
func GetProducts() Products {
	return productList
}

// AddProduct adds a product with p data
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

// DeleteProduct removes a product with given id
func DeleteProduct(id int) error {
	_, idx, err := findProductByID(id)

	if err != nil {
		return err
	}
	productList = append(productList[:idx], productList[idx+1:]...)
	return nil
}

// UpdateProduct updates a product with given id
func UpdateProduct(id int, p *Product) error {
	_, idx, err := findProductByID(id)

	if err != nil {
		return err
	}

	p.ID = id
	productList[idx] = p
	return nil
}

// ErrProductNotFound used in http return errors
var ErrProductNotFound = fmt.Errorf("Product not found")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

func findProductByID(id int) (*Product, int, error) {
	for i, prod := range productList {
		if prod.ID == id {
			return prod, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}
func getNextID() int {
	lastID := productList[len(productList)-1].ID
	return lastID+1
}

// findIndex finds the index of a product in the database
// returns -1 when no product can be found
func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

// GetProductByID returns a single product which matches the id from the
// database.
// If a product is not found this function returns a ProductNotFound error
func GetProductByID(id int) (*Product, error) {
	i := findIndexByProductID(id)
	if id == -1 {
		return nil, ErrProductNotFound
	}

	return productList[i], nil
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "hfg-dsd-fff",
	},
	&Product{
		ID:          2,
		Name:        "Esspresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd-hhh-aaa",
	},
}