package product

import "context"

type ProductRepository interface {
	Save(ctx context.Context, product *Product) error
	FindByID(ctx context.Context,id string) (*Product, error)
	FindByIDs(ctx context.Context,id []string) ([]Product, error)
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, product *Product) error
}