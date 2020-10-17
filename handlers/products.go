// Package classification of Product API 
//
// Documentation for Product API 
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes: 
// - application/json 
//
// Products:
// - application/json
// swagger:meta
package handlers

import (
	"github.com/cap-diego/microservices/data"
	"log"
	"fmt"
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
