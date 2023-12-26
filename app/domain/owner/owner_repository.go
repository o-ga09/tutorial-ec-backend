package owner

import "context"

type OwnerRepository interface {
	Save(ctx context.Context) error
	FindById(ctx context.Context, id string) (*Owner, error)
}