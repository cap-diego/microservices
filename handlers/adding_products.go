package handlers

import (
	"net/http"
	"github.com/cap-diego/microservices/data"
)

//swager:route POST /products products createProduct
//
// responses:
//	200: productResponse
//	422: errorValidation
//	501: errorResponse

// Create handles post request for creating a new product
func (prods *Products) Create(rw http.ResponseWriter, req *http.Request) {
	newProd := req.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&newProd)
}
