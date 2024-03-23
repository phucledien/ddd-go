package tavern

import (
	"testing"

	"github.com/google/uuid"

	"github.com/phucledien/tavern/domain/product"
	"github.com/phucledien/tavern/services/order"
)

func Test_Tavern(t *testing.T) {
  // Create orderService
  products := init_products(t)
  os, err := order.NewOrderService(
    order.WithMongoCustomerRepository("mongodb://localhost:27017"),
    order.WithMemoryProductRepository(products),
  )
  if err != nil {
    t.Error(err)
  }

  ts, err := NewTavernService(WithOrderService(os))
  if err != nil {
    t.Error(err)
  }

  cust, err := os.AddCustomer("Percy")
  order := []uuid.UUID{
    products[0].GetID(),
  }
  // Execute Order
  err = ts.Order(cust.GetID(), order)
  if err != nil {
    t.Error(err)
  }
}

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
