package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"

	"github.com/phucledien/tavern/domain/customer"
)

type MemoryReposiory struct {
	customers map[uuid.UUID]customer.Customer
	sync.Mutex
}

func New() *MemoryReposiory {
	return &MemoryReposiory{
		customers: make(map[uuid.UUID]customer.Customer),
	}
}

func (mr *MemoryReposiory) Get(id uuid.UUID) (customer.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}

	return customer.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryReposiory) Add(c customer.Customer) error {
	if mr.customers == nil {
		// Safety check if customers is not create, shoudn't happen if using the factory, but you never know
		mr.Lock()
		mr.customers = make(map[uuid.UUID]customer.Customer)
		mr.Unlock()
	}
	// Make sure Customer isn't already in the repository
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFailedToAddCustomer)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}

func (mr *MemoryReposiory) Update(c customer.Customer) error {
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}
