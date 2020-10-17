package handlers

import (
	"net/http"
	"github.com/cap-diego/microservices/data"
)

// AddProduct http handler of post
func (prods *Products) AddProduct(rw http.ResponseWriter, req *http.Request) {
	newProd := req.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&newProd)
}
