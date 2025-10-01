package service

import (
	"backend-ecommerce/internal/application/entity"
	"backend-ecommerce/internal/application/repository"
)

type orderService struct {
	repo          repository.OrderRepository
	cartService   CartService
}

// NewOrderService creates a new order service
func NewOrderService(
	repo repository.OrderRepository,
	cartService CartService,
) OrderService {
	return &orderService{
		repo:        repo,
		cartService: cartService,
	}
}

func (s *orderService) CreateOrder(order entity.Order) (*entity.Order, error) {
	// Add any business logic here before creating
	// For example, validate cart, calculate totals, etc.
	return s.repo.Create(order)
}

func (s *orderService) GetOrderByID(id string) (*entity.Order, error) {
	return s.repo.FindByID(id)
}

func (s *orderService) GetUserOrders(userID string, page, pageSize int) ([]entity.Order, int64, error) {
	return s.repo.FindByUserID(userID, page, pageSize)
}

func (s *orderService) UpdateOrderStatus(id string, status string) (*entity.Order, error) {
	order, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	order.Status = status
	return s.repo.Update(*order)
}

func (s *orderService) CancelOrder(id string) error {
	order, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	// Only allow canceling orders that are in a cancelable state
	if order.Status != "pending" && order.Status != "processing" {
		return ErrCannotCancelOrder
	}

	order.Status = "cancelled"
	_, err = s.repo.Update(*order)
	return err
}

// Error variables
var (
	ErrCannotCancelOrder = NewServiceError("cannot cancel order in current state")
)
