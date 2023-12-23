package order

// ========================================== //
// 注文処理                                    //
// ・ユーザーはカート内の商品を注⽂できる           //
// ・在庫を超える量の商品を注⽂することは不可能      //
// ・注⽂には最低 1 つの商品が含まれている必要がある //
// ・削除された商品を注⽂することはできない         //
// ・注⽂すると、その履歴を保存する                //
// ========================================== //

import (
	"context"
	"time"

	"github.com/google/uuid"
	cartDomain "github.com/o-ga09/tutorial-go-fr/app/domain/cart"
	"github.com/o-ga09/tutorial-go-fr/pkg/strings"
)

type Order struct {
	id string
	userID string
	totalAmount int64
	product OrderProducts
	orderAt time.Time
}

type OrderProducts []OrderProduct

type OrderProduct struct {
	productID string
	price int64
	quantity int
}

type OrderDomainService interface {
	OrderProducts(ctx context.Context,cart *cartDomain.Cart, now time.Time) (string, error)
}

func NewOrder(userID string, totalAmount int64, products []OrderProduct, now time.Time) (*Order, error) {
	id := strings.RemoveHyphen(uuid.NewString())
	return newOrder(id,userID,totalAmount,products,now)
}

func Reconstract(id string,userID string, totalAmount int64, products []OrderProduct, now time.Time) (*Order, error) {
	return newOrder(id,userID,totalAmount,products,now)
}

func newOrder(id string,userID string, totalAmount int64, products []OrderProduct, now time.Time) (*Order, error) {
	return &Order{
		id: id,
		userID: userID,
		totalAmount: totalAmount,
		product: products,
		orderAt: now,
	}, nil
}

func(o *Order) UserID() string {return o.userID}
func(o *Order) ID() string {return o.id}
func(o *Order) Totalamount() int64 {return o.totalAmount}
func(o *Order) Products() []OrderProduct {return o.product}
func(o *Order) OrderAt() time.Time {return o.orderAt}
func(o *Order) ProductIDs() []string {
	var productIDs [] string
	for _, p := range o.product {
		productIDs = append(productIDs,p.productID)
	}
	return productIDs
}

func(o OrderProducts) ProductIDs() []string {
	var productIDs []string
	for _, p := range o {
		productIDs = append(productIDs, p.productID)
	}
	return productIDs
}

func(o OrderProducts) TotalAmount() int64 {
	var totalAmount int64
	for _, p := range o {
		totalAmount += p.price * int64(p.quantity)
	}

	return totalAmount
}

func NewOrderProduct(productID string, price int64, quantity int) (*OrderProduct, error) {
	return &OrderProduct{
		productID: productID,
		price: price,
		quantity: quantity,
	}, nil
}

func(o *OrderProduct) ProductID() string {return o.productID}
func(o *OrderProduct) Price() int64 {return o.price}
func(o *OrderProduct) Quantity() int {return o.quantity}