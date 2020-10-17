package handlers

import (
	"net/http"
	"github.com/cap-diego/microservices/data"
)

// GetProducts http handler of GET
func (prods *Products) GetProducts(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
