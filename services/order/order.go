package order

import (
	"context"
	"log"

	"github.com/google/uuid"

	"github.com/phucledien/tavern/domain/customer"
	"github.com/phucledien/tavern/domain/customer/memory"
	"github.com/phucledien/tavern/domain/customer/mongo"
	"github.com/phucledien/tavern/domain/product"
	prodMemory "github.com/phucledien/tavern/domain/product/memory"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.Repository
	products  product.Repository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithCustomerRepository(cr customer.Repository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMemoryProductRepository(products []product.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodMemory.New()
		for _, product := range products {
			err := pr.Add(product)
			if err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

func WithMongoCustomerRepository(connectionString string) OrderConfiguration {
  return func(os *OrderService) error { 
    cr, err := mongo.New(context.Background(), connectionString)
    if err != nil {
      return err
    }
    os.customers = cr
    return nil
  }
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	var products []product.Product
	var total float64
	for _, id := range productIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		total += p.GetPrice()
	}

	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))
	return total, nil
}

func (o *OrderService) AddCustomer(name string) (customer.Customer, error) {
  cust, err := customer.New(name)
  if err != nil {
    return customer.Customer{}, err
  }

  err = o.customers.Add(cust)
  if err != nil {
    return customer.Customer{}, err
  }

  return cust, nil
}


