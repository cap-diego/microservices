package handlers

import (
	"github.com/gorilla/mux"
	"github.com/cap-diego/microservices/data"
	"log"
	"strconv"
	"net/http"
	"context"
)

// Products rest resource, implements ServeHTTP
type Products struct {
	l *log.Logger
}

type KeyProduct struct{}

// NewProducts rest resource for products
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (prods *Products) GetProducts(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (prods *Products) AddProduct(rw http.ResponseWriter, req *http.Request) {
	newProd := req.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&newProd)
}

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

func (prods Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	f := func(rw http.ResponseWriter, req *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(req.Body)
		if err != nil {
			prods.l.Println("[ERROR] desarializing product",err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		// Add product to the context 
		ctx := context.WithValue(req.Context(), KeyProduct{}, prod)
		req = req.WithContext(ctx)
		prods.l.Printf("CONTEXTO %#v", ctx)
		// Call the next handler 
		next.ServeHTTP(rw, req)
	}
	return http.HandlerFunc(f)
}





//With marshal
// d, err := json.Marshal(lp)
// if err != nil {
// 	http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
// }
// rw.Write(d)

