package service

import (
	"backend-ecommerce/internal/application/entity"
	"backend-ecommerce/internal/application/repository"
)

type paymentService struct {
	repo repository.PaymentRepository
}

// NewPaymentService creates a new payment service
func NewPaymentService(repo repository.PaymentRepository) PaymentService {
	return &paymentService{
		repo: repo,
	}
}

func (s *paymentService) ProcessPayment(payment entity.Payment) (*entity.Payment, error) {
	// Add payment processing logic here (e.g., integrate with payment gateway)
	// For now, we'll just save it with a success status
	payment.Status = "succeeded"
	return s.repo.Create(payment)
}

func (s *paymentService) GetPaymentByID(id string) (*entity.Payment, error) {
	return s.repo.FindByID(id)
}

func (s *paymentService) GetPaymentsByOrderID(orderID string) ([]entity.Payment, error) {
	return s.repo.FindByOrderID(orderID)
}

func (s *paymentService) RefundPayment(paymentID string, amount float64) error {
	payment, err := s.repo.FindByID(paymentID)
	if err != nil {
		return err
	}

	// Add refund logic here (e.g., call payment gateway's refund API)
	// For now, we'll just update the status
	payment.Status = "refunded"
	_, err = s.repo.Update(*payment)
	return err
}
