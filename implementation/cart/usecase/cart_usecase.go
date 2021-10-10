package usecase

import (
	"fmt"
	"github.com/huf0813/gemastik_api_backend_supabase/domain"
)

type CartUseCase struct {
	CartRepositorySupabase domain.CartRepository
}

func NewCartUseCase(c domain.CartRepository) domain.CartUseCase {
	return &CartUseCase{CartRepositorySupabase: c}
}

func (cart *CartUseCase) DeleteProductFromCart(userID, productID, cartID int64) error {
	return cart.CartRepositorySupabase.DeleteProductFromCart(userID, productID, cartID)
}

func (cart *CartUseCase) AddProductToCart(c *domain.AddProductToCartRequest, userID int64) error {
	value := map[string]string{
		"user_id":    fmt.Sprintf("%d", userID),
		"product_id": fmt.Sprintf("%d", c.ProductID),
		"quantity":   fmt.Sprintf("%d", c.Quantity),
	}

	if err := cart.CartRepositorySupabase.AddProductToCart(value); err != nil {
		return err
	}

	return nil
}
