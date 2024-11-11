package services

import (
	"errors"
	"golang-beginner-chap24/collections"
	"golang-beginner-chap24/repositories"
)

type PaymentMethodService struct {
	PaymentMethodRepo *repositories.PaymentMethodRepository
}

func NewPaymentMethodService(paymentMethodRepo *repositories.PaymentMethodRepository) *PaymentMethodService {
	return &PaymentMethodService{PaymentMethodRepo: paymentMethodRepo}
}

func (s *PaymentMethodService) GetAllPaymentMethods() ([]collections.PaymentMethod, error) {
	return s.PaymentMethodRepo.GetAll()
}

func (s *PaymentMethodService) GetPaymentMethodById(id int) (*collections.PaymentMethod, error) {
	if id <= 0 {
		return nil, errors.New("invalid payment method id")
	}
	paymentMethod, err := s.PaymentMethodRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &paymentMethod, nil
}

func (s *PaymentMethodService) CreatePaymentMethod(paymentMethodInput collections.PaymentMethod) error {
	if paymentMethodInput.Name == "" {
		return errors.New("payment method name is required")
	}

	return s.PaymentMethodRepo.Create(paymentMethodInput)
}

func (s *PaymentMethodService) UpdatePaymentMethod(id int, paymentMethodInput collections.PaymentMethod) error {
	if id <= 0 {
		return errors.New("invalid payment method id")
	}

	return s.PaymentMethodRepo.Update(id, paymentMethodInput)
}

func (s *PaymentMethodService) DeletePaymentMethod(id int) error {
	if id <= 0 {
		return errors.New("invalid payment method id")
	}
	return s.PaymentMethodRepo.Delete(id)
}
