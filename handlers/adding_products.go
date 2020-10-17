package handlers

import (
	"net/http"
	"github.com/cap-diego/microservices/data"
)

// swagger:route POST /products products createProduct
// Create a new product
// responses:
//	200: productResponse
//	422: errorValidation
//	501: errorResponse

// Create handles post request for creating a new product
func (prods *Products) Create(rw http.ResponseWriter, req *http.Request) {
	newProd := req.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&newProd)
}
