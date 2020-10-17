package handlers

import (
	"github.com/gorilla/mux"
	"github.com/cap-diego/microservices/data"
	"log"
	"fmt"
	"strconv"
	"net/http"
	"context"
)

// Products rest resource, implements ServeHTTP
type Products struct {
	l *log.Logger
}

// KeyProduct used for request.context 
type KeyProduct struct{}

// NewProducts rest resource for products
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// GetProducts http handler of GET
func (prods *Products) GetProducts(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// AddProduct http handler of post
func (prods *Products) AddProduct(rw http.ResponseWriter, req *http.Request) {
	newProd := req.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&newProd)
}

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

//MiddlewareProductValidation Validate body of request and save product to context if valid
func (prods Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	f := func(rw http.ResponseWriter, req *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(req.Body)
		if err != nil {
			prods.l.Println("[ERROR] desarializing product",err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		// Validate products

		err = prod.Validate()
		if err != nil {
			prods.l.Println("[ERROR] validating product", err)
			http.Error(
				rw, 
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest)
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

