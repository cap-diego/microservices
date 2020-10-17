package main

import (
	"testing"
	"fmt"
	"github.com/cap-diego/microservices/client/client"
	"github.com/cap-diego/microservices/client/client/products"

)


func TestOurClient(t *testing.T) {
	config := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, config)
	params := products.NewListProductsParams()
	prods, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Res: %#v", prods.GetPayload()[0])
	t.Fail()
	// fmt.Println(prods)
}