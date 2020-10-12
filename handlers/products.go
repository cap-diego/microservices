package handlers

import (
	"log"
	"net/http"
	"github.com/cap-diego/microservices/data"
)

// Products rest resource
type Products struct {
	l *log.Logger
}

// NewProducts rest resource for products
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		p.getProducts(rw, req)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err!= nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

//With marshal
// d, err := json.Marshal(lp)
// if err != nil {
// 	http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
// }
// rw.Write(d)