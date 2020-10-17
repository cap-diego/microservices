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
	"strconv"
	"github.com/gorilla/mux"
)

// KeyProduct used for request.context 
type KeyProduct struct{}

// Products rest resource, implements ServeHTTP
type Products struct {
	l *log.Logger
	v *data.Validation
}

// NewProducts rest resource for products
// func NewProducts(l *log.Logger) *Products {
// 	return &Products{l}
// }

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// NewProducts returns a new products handler with the given logger
func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}
// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getProductID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}