package data

import (
	"testing"
)
func TestChecksProductWithNoNameIsNotValid(t *testing.T) {
	p := &Product{Price: 1, SKU: "abc-hgs-bbb"}

	err := p.Validate()

	if err == nil {
		t.Fatal(err)
	}
}

func TestChecksProductWithNameIsValid(t *testing.T) {
	p := &Product{Name: "Donnut", Price: 10, SKU: "abc-hgs-bbb"}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

func TestChecksProductWithPriceZeroIsInvalid(t *testing.T) {
	p := &Product{Name: "Donnut", Price: 0, SKU: "abc-hgs-bbb"}
	err := p.Validate()
	if err == nil {
		t.Fatal(err)
	}
}

func TestValidateSKU(t *testing.T) {
	p := &Product{Name: "Donnut", Price: 10, SKU: "invalidsku"}
	err := p.Validate()
	if err == nil {
		t.Fatal(err)
	}
}

func TestValidaRightSKU(t *testing.T) {
	p := &Product{Name: "Donnut", Price: 10, SKU: "abc-hgs-bbb"}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}