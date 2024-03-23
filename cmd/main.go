package main

import (
	"github.com/google/uuid"

	"github.com/phucledien/tavern/domain/product"
	"github.com/phucledien/tavern/services/order"
	"github.com/phucledien/tavern/services/tavern"
)

func main() {
  products := productInventory()

  os, err := order.NewOrderService(
    order.WithMongoCustomerRepository("mongodb://localhost:27017"),
    order.WithMemoryProductRepository(products),
  )
  if err != nil {
    panic(err)
  }

  ts, err := tavern.NewTavernService(
    tavern.WithOrderService(os),
  )
  cust, err := os.AddCustomer("John Doe")
  if err != nil {
    panic(err)
  }
  orders := []uuid.UUID {
    products[0].GetID(),
  }

  err = ts.Order(cust.GetID(), orders)
  if err != nil {
    panic(err)
  }
}


func productInventory() []product.Product {
  beer , err := product.New("Beer", "Healthy Beverage", 1.99)
  if err != nil {
    panic(err)
  }
  peenuts, err := product.New("Peenuts", "Healthy Snack", 0.99)
  if err != nil {
    panic(err)
  }
  wine, err := product.New("Wine", "Healthy Beverage", 0.99)
  if err != nil {
    panic(err)
  }
  products := []product.Product{
    beer, peenuts, wine,
  }
  return products
}
