// Services aggregates multiple repositories
package services

import (
	"log"

	"github.com/arturbaccarin/programmingpercy-go-videos/ddd-go/aggregate"
	"github.com/arturbaccarin/programmingpercy-go-videos/ddd-go/domain/customer"
	"github.com/arturbaccarin/programmingpercy-go-videos/ddd-go/domain/customer/memory"
	"github.com/arturbaccarin/programmingpercy-go-videos/ddd-go/domain/product"
	prodmem "github.com/arturbaccarin/programmingpercy-go-videos/ddd-go/domain/product/memory"
	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository

	// billing billing.Service
}

// services can hold multiple repositories
// services are allowed to hold other services

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	// Loop through all the cfgs and apply them
	for _, cfg := range cfgs {
		err := cfg(os)

		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

/*
NewOrderService(
	WithMemoryCustomerRepository(),
	WithLogging("debug"),
	WithTracing()
)
*/

// WithCustomerRepository applies a customer repository to the OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.New()

		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}
		}

		os.products = pr
		return nil
	}
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	// Fetch the customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	// Get each Product
	var products []aggregate.Product
	var total float64

	for _, id := range productIDs {
		p, err := o.products.GetByID(id)

		if err != nil {
			return 0, err
		}

		products = append(products, p)
		total += p.GetPrice()
	}

	log.Printf("Customer: %s has ordered %d products for a total of $%.2f", c.GetName(), len(products), total)
	return total, nil
}
