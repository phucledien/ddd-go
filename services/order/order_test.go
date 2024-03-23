package order

import (
	"testing"

	"github.com/google/uuid"
	"github.com/phucledien/tavern/domain/product"
	"github.com/phucledien/tavern/domain/customer"
)

func init_products(t *testing.T) []product.Product {
  beer , err := product.New("Beer", "Healthy Beverage", 1.99)
  if err != nil {
    t.Error(err)
  }
  peenuts, err := product.New("Peenuts", "Healthy Snack", 0.99)
  if err != nil {
    t.Error(err)
  }
  wine, err := product.New("Wine", "Healthy Beverage", 0.99)
  if err != nil {
    t.Error(err)
  }
  products := []product.Product{
    beer, peenuts, wine,
  }
  return products
}

func TestOrder_NewOrderService(t *testing.T) {
  products := init_products(t)
  os, err := NewOrderService(
    WithMemoryCustomerRepository(),
    WithMemoryProductRepository(products),
  )
  if err != nil {
    t.Error(err)
  }

  // Add Customer
  cust, err := customer.New("Percy")
  if err != nil {
    t.Error(err)
  }

  err = os.customers.Add(cust)
  if err != nil {
    t.Error(err)
  }

  // Perform Order for one beer
  order := []uuid.UUID{
    products[0].GetID(),
  }
  _, err = os.CreateOrder(cust.GetID(), order)
  if err != nil {
    t.Error(err)
  }
}
