package services

import (
	"testing"

	"github.com/arturbaccarin/programmingpercy-go-videos/ddd-go/aggregate"
	"github.com/google/uuid"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "Bottle of beer", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	peenuts, err := aggregate.NewProduct("Peenuts", "Bag of Peenuts", 1.99)
	if err != nil {
		t.Fatal(err)
	}

	wine, err := aggregate.NewProduct("Wine", "Glass of wine", 9.99)
	if err != nil {
		t.Fatal(err)
	}

	return []aggregate.Product{
		beer, peenuts, wine,
	}
}

func TestOrder_NewOrderService(t *testing.T) {
	products := init_products(t)

	// loading repositories
	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("John")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}
