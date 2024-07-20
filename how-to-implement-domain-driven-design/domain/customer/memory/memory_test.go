package memory

import (
	"errors"
	"testing"

	"github.com/arturbaccarin/programmingpercy-go-videos/ddd-go/aggregate"
	"github.com/arturbaccarin/programmingpercy-go-videos/ddd-go/domain/customer"
	"github.com/google/uuid"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	cust, err := aggregate.NewCustomer("john")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetID()

	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name:        "no customer by id",
			id:          uuid.MustParse("7faa1d3b-5030-4e2e-993d-80f9782fb5f1"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "customer by id",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
