package service

import (
	"backend-ecommerce/internal/application/entity"
	"backend-ecommerce/internal/application/repository"
)

type cartService struct {
	repo repository.CartRepository
}

// NewCartService creates a new cart service
func NewCartService(repo repository.CartRepository) CartService {
	return &cartService{
		repo: repo,
	}
}

func (s *cartService) GetOrCreateCart(userID *string, guestToken string) (*entity.Cart, error) {
	var cart *entity.Cart
	var err error

	// If user is authenticated, try to get their cart
	if userID != nil && *userID != "" {
		cart, err = s.repo.FindByUserID(*userID)
		if err == nil && cart != nil {
			return cart, nil
		}

		// If no cart found for user, create a new one
		cart = &entity.Cart{
			UserID: userID,
		}
	} else if guestToken != "" {
		// Try to get guest cart
		cart, err = s.repo.FindByGuestToken(guestToken)
		if err == nil && cart != nil {
			return cart, nil
		}

		// If no cart found for guest, create a new one
		cart = &entity.Cart{
			GuestToken: &guestToken,
		}
	} else {
		// Create a new guest cart
		cart = &entity.Cart{}
	}

	return s.repo.Create(*cart)
}

func (s *cartService) GetCartByID(id string) (*entity.Cart, error) {
	return s.repo.FindByID(id)
}

func (s *cartService) AddCartItem(cartID string, item entity.CartItem) (*entity.Cart, error) {
	// Get the cart
	cart, err := s.repo.FindByID(cartID)
	if err != nil {
		return nil, err
	}

	// Check if item already exists in cart
	for i, existingItem := range cart.Items {
		if existingItem.ProductID == item.ProductID {
			// Update quantity if item exists
			cart.Items[i].Quantity += item.Quantity
			return s.repo.Update(*cart)
		}
	}

	// Add new item to cart
	cart.Items = append(cart.Items, item)
	return s.repo.Update(*cart)
}

func (s *cartService) UpdateCartItem(cartID, itemID string, quantity int) (*entity.Cart, error) {
	// Get the cart
	cart, err := s.repo.FindByID(cartID)
	if err != nil {
		return nil, err
	}

	// Find and update the item
	for i, item := range cart.Items {
		if item.ID == itemID {
			if quantity <= 0 {
				// Remove item if quantity is 0 or less
				cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			} else {
				// Update quantity
				cart.Items[i].Quantity = quantity
			}
			return s.repo.Update(*cart)
		}
	}

	return cart, nil
}

func (s *cartService) RemoveCartItem(cartID, itemID string) error {
	cart, err := s.repo.FindByID(cartID)
	if err != nil {
		return err
	}

	// Find and remove the item
	for i, item := range cart.Items {
		if item.ID == itemID {
			cart.Items = append(cart.Items[:i], cart.Items[i+1:]...)
			_, err := s.repo.Update(*cart)
			return err
		}
	}

	return nil
}

func (s *cartService) ClearCart(cartID string) error {
	cart, err := s.repo.FindByID(cartID)
	if err != nil {
		return err
	}

	// Clear all items
	cart.Items = []entity.CartItem{}
	_, err = s.repo.Update(*cart)
	return err
}
