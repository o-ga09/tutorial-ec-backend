package order

import "context"

type OrderRepository interface {
	Save(cyx context.Context, order *Order) error
}