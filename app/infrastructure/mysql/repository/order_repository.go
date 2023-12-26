package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/o-ga09/tutorial-ec-backend/app/domain/order"
	"github.com/o-ga09/tutorial-ec-backend/pkg/strings"
	"gorm.io/gorm"
)

type OrderRepository struct{
	conn *gorm.DB
}

type Order struct {
	id string
	userID string
	totalAmount int64
	orderAt time.Time
}

type OrderProduct struct {
	id string
	orderID string
	productID string
	price int64
	quantity int
}

// Save implements order.OrderRepository.
func (o *OrderRepository) Save(cyx context.Context, order *order.Order) error {
	repoOrder := Order{
		id: order.ID(),
		userID: order.ID(),
		totalAmount: order.Totalamount(),
		orderAt: order.OrderAt(),
	}
	o.conn.Create(repoOrder)

	op := order.Products()
	for _, p := range op {
		id := strings.RemoveHyphen(uuid.New().String())
		repoOR := OrderProduct{
			id: id,
			productID: p.ProductID(),
			orderID: order.ID(),
			price: p.Price(),
			quantity: p.Quantity(),
		}
		o.conn.Create(repoOR)
	}
	return nil
}

func NewOrderRepository(conn *gorm.DB) order.OrderRepository {
	return &OrderRepository{conn: conn}
}
