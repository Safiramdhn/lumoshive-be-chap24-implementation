package services

import (
	"errors"
	"golang-beginner-chap24/collections"
	"golang-beginner-chap24/repositories"
	"reflect"
)

type OrderService struct {
	OrderRepo *repositories.OrderRepository
}

func NewOrderService(orderRepo *repositories.OrderRepository) *OrderService {
	return &OrderService{OrderRepo: orderRepo}
}

func (s *OrderService) CreateOrder(orderInput collections.Order) error {
	if reflect.DeepEqual(orderInput.ShippingAddress, collections.Address{}) {
		return errors.New("customer address cannot be empty")
	}
	if len(orderInput.OrderItems) == 0 {
		return errors.New("order items cannot be empty")
	}

	err := s.OrderRepo.Create(orderInput)
	if err != nil {
		return err
	}
	return nil
}
