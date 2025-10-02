package service

import (
	"backend-ecommerce/internal/application/entity"
)

type orderService struct {
}


func (s *orderService) CreateOrder(order entity.Order) (*entity.Order, error) {

	return nil, nil
}

func (s *orderService) GetOrderByID(id string) (*entity.Order, error) {
	return nil, nil
}

func (s *orderService) GetUserOrders(userID string, page, pageSize int) ([]entity.Order, int64, error) {
	return []entity.Order{}, 0, nil
}

func (s *orderService) UpdateOrderStatus(id string, status string) (*entity.Order, error) {
	return nil, nil
}

func (s *orderService) CancelOrder(id string) error {
	return nil
}