package handlers

import (
	"github.com/cap-diego/microservices/data"
	"log"
	"regexp"
	"strconv"
	"net/http"
)

// Products rest resource, implements ServeHTTP
type Products struct {
	l *log.Logger
}

// NewProducts rest resource for products
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

//ServeHTTP is the main entry point for the products. Satisfies http.Handler interface
func (prods *Products) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		prods.getProducts(rw, req)
		return
	}

	if req.Method == http.MethodPost {
		prods.addProduct(rw, req)
		return
	}

	if req.Method == http.MethodPut {
		//Expects id in the url
		reg := regexp.MustCompile(`/([0-9]+)`)
		res := reg.FindAllStringSubmatch(req.URL.Path, -1)
		if len(res) != 1{
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		if len(res[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := res[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		prods.updateProducts(id, rw, req)
		return 
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (prods *Products) getProducts(rw http.ResponseWriter, h *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (prods *Products) addProduct(rw http.ResponseWriter, req *http.Request) {
	newProd := &data.Product{}
	err := newProd.FromJSON(req.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	data.AddProduct(newProd)
}

func (prods *Products) updateProducts(id int, rw http.ResponseWriter, req *http.Request) {
	prod := &data.Product{}
	err := prod.FromJSON(req.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	} 
}

//With marshal
// d, err := json.Marshal(lp)
// if err != nil {
// 	http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
// }
// rw.Write(d)
