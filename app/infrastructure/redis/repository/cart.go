package repository

import (
	"context"

	cartDomain "github.com/o-ga09/tutorial-ec-backend/app/domain/cart"
)

type cartRepository struct{}

type cartProduct struct {
	ProductID string `json:"product_id,omitempty"`
	Quantity  int    `json:"quantity,omitempty"`
}

// FindByID implements cart.CartRepository.
func (*cartRepository) FindByID(ctx context.Context, userID string) (*cartDomain.Cart, error) {
	panic("unimplemented")
}

// Save implements cart.CartRepository.
func (*cartRepository) Save(ctx context.Context, cart *cartDomain.Cart) error {
	panic("unimplemented")
}

func NewCartRepository() cartDomain.CartRepository {
	return &cartRepository{}
}
