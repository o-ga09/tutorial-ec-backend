package repository

import (
	"context"

	"github.com/o-ga09/tutorial-ec-backend/app/domain/order"
	"github.com/o-ga09/tutorial-ec-backend/app/infrastructure/mysql/schema"
	"gorm.io/gorm"
)

type OrderRepository struct{
	conn *gorm.DB
}

// Save implements order.OrderRepository.
func (o *OrderRepository) Save(cyx context.Context, order *order.Order) error {
	repoOrder := schema.Order{
		UserID: order.ID(),
		TotalAmount: order.Totalamount(),
		OrderAt: order.OrderAt(),
	}
	o.conn.Create(repoOrder)

	op := order.Products()
	for _, p := range op {
		repoOR := schema.OrderProduct{
			ProductID: p.ProductID(),
			OrderID: order.ID(),
			Price: p.Price(),
			Quantity: p.Quantity(),
		}
		o.conn.Create(repoOR)
	}
	return nil
}

func NewOrderRepository(conn *gorm.DB) order.OrderRepository {
	return &OrderRepository{conn: conn}
}
