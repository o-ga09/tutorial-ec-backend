package cart

import "context"

//go:generate moq -out CartRepository_mock.go . CartRepository
type CartRepository interface {
	FindByID(ctx context.Context, userID string) (*Cart, error)
	Save(ctx context.Context, cart *Cart) error
}