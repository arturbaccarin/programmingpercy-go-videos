// Services aggregates multiple repositories
package services

import "github.com/arturbaccarin/programmingpercy-go-videos/ddd-go/domain/customer"

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
}

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
