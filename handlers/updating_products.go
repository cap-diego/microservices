package handlers

import (
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/cap-diego/microservices/data"
)

// UpdateProducts http handler of put
func (prods *Products) UpdateProducts(rw http.ResponseWriter, req *http.Request) {
	requestVars := mux.Vars(req)
	id, err := strconv.Atoi(requestVars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert ID from url", http.StatusBadRequest)
		return
	}

	prod := req.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	} 
}