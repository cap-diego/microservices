package handlers

import (
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/cap-diego/microservices/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Deletes a product from the db
// responses:
//	201: noContent

//DeleteProduct delete the product from the db 
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.l.Println("Handle DELETE product", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
}