package tavern

import (
	"log"

	"github.com/google/uuid"

  "github.com/phucledien/tavern/services/order"
)

type TavernConfiguration func(ts *TavernService) error

type TavernService struct {
  // Orderservice is used to handle orders
  OrderService *order.OrderService

  // BillingService is used to handle billing
  // This is up to you to implement
  BillingService interface{}
}

func NewTavernService(cfgs ...TavernConfiguration) (*TavernService, error) {
  ts := &TavernService{}

  for _, cfg := range cfgs {
    if err := cfg(ts); err != nil {
      return nil, err
    }
  }

  return ts, nil
}

func WithOrderService(os *order.OrderService) TavernConfiguration {
  return func(ts *TavernService) error {
    ts.OrderService = os
    return nil
  }
}

func (ts *TavernService) Order(customer uuid.UUID, products []uuid.UUID) error {
  total, err := ts.OrderService.CreateOrder(customer, products)
  if err != nil {
    return err
  }
  log.Printf("Bill the customer: %0.0f", total)

  // Bill the customer
  // err = t.BillingService.Bill(customer, total)
  return nil
}
