package service

import (
	"backend-ecommerce/internal/application/dto"
	"backend-ecommerce/internal/application/entity"
)

type cartService struct {
}

func (s *cartService) GetOrCreateCart(userID *string, guestToken string) dto.ResponseDto {
	return dto.ResponseDto{}
}

func (s *cartService) GetCartByID(id string) (*entity.Cart, error) {
	return nil, nil
}

func (s *cartService) AddCartItem(cartID string, item entity.CartItem) (*entity.Cart, error) {
	return nil, nil
}

func (s *cartService) UpdateCartItem(cartID, itemID string, quantity int) (*entity.Cart, error) {
	return nil, nil
}

func (s *cartService) RemoveCartItem(cartID, itemID string) error {
	return nil
}

func (s *cartService) ClearCart(cartID string) error {
	return nil
}