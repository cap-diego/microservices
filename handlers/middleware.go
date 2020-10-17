package handlers

import (
	"context"
	"github.com/cap-diego/microservices/data"
	"net/http"
)

// MiddlewareProductValidation Validate body of request and save product to context if valid
func (prods *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	f := func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		prod := data.Product{}
		err := data.FromJSON(&prod, req.Body)
		if err != nil {
			prods.l.Println("[ERROR] desarializing product",err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		// Validate products

		errors := prods.v.Validate(prod)
		if len(errors) != 0 {
			prods.l.Println("[ERROR] validating product", err)
			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errors.Errors()}, rw)
			return			
		}

		// Add product to the context 
		ctx := context.WithValue(req.Context(), KeyProduct{}, prod)
		req = req.WithContext(ctx)
		// Call the next handler 
		next.ServeHTTP(rw, req)
	}
	return http.HandlerFunc(f)
}