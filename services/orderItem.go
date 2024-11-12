package services

import "golang-beginner-chap24/repositories"

type OrderItemService struct {
	OrderItemRepo *repositories.OrderItemRepository
}

func NewOrderItemService(orderItemRepo *repositories.OrderItemRepository) *OrderItemService {
	return &OrderItemService{OrderItemRepo: orderItemRepo}
}

func (s *OrderItemService) GetRatingAverage() (float64, error) {
	return s.OrderItemRepo.CalculateAverageRating()
}
