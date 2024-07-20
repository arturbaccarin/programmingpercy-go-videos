// package aggregate holds our aggrets that combine entities and valueobjects
// in a full object
package aggregate

import (
	"errors"

	"github.com/arturbaccarin/programmingpercy-go-videos/ddd-go/entity"
	"github.com/arturbaccarin/programmingpercy-go-videos/ddd-go/valueobject"
	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("a customer must have a name")
)

type Customer struct {
	// person is the root entity of customer
	// which means person.ID is the main identifier for the customer
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction
}

// NewCustomer is a factory to create a new customer aggregate
// it will validate that the name is not empty
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: []valueobject.Transaction{},
	}, nil
}
