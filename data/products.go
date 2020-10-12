package data

import (
	"time"
	"io"
	"fmt"
	"encoding/json"
)

type Product struct {
	ID			int 	`json:"id"`
	Name		string  `json:"name"`
	Desc		string  `json:"description"`
	Price		float32 `json:"price"`
	SKU			string	`json:"sku"`
	CreatedOn	string  `json:"-"`
	DeletedOn	string  `json:"-"`
	UpdatedOn	string	`json:"-"`
}

type Products []*Product


func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *Product) error {
	_, idx, err := findProductById(id)

	if err != nil {
		return err
	}

	p.ID = id
	productList[idx] = p
	return nil
}

var ErrProductNotFound = fmt.Errorf("Product not found")

func findProductById(id int) (*Product, int, error) {
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

var productList = []*Product{
	&Product{
		ID:			1,
		Name:		"Latte",
		Desc:		"Frothy milky coffee",
		SKU: 		"--",
		Price:		2.5,
		CreatedOn:	time.Now().UTC().String(),
		UpdatedOn:  time.Now().UTC().String(),
	},
	&Product{
		ID:			2,
		Name:		"Espresso",
		Desc:		"Short and strong coffe without milk",
		SKU: 		"--",
		Price:		1.5,
		CreatedOn:	time.Now().UTC().String(),
		UpdatedOn:  time.Now().UTC().String(),
	},
}