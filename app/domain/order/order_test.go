package order

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"github.com/o-ga09/tutorial-go-fr/pkg/strings"
)

func TestNewOrder(t *testing.T) {
	userID := strings.RemoveHyphen(uuid.New().String())
	productID1 := strings.RemoveHyphen(uuid.New().String())
	productID2 := strings.RemoveHyphen(uuid.New().String())
	orderAt := time.Now()
	type args struct {
		totalAmount int64
		products    []OrderProduct
	}
	tests := []struct {
		name    string
		args    args
		want    *Order
		wantErr bool
	}{
		{name: "正常系",args: args{totalAmount: 100, products: []OrderProduct{{productID: productID1,quantity: 1},{productID: productID2,quantity: 2}}},want: &Order{totalAmount: 100, userID: userID,product: []OrderProduct{{productID: productID1,quantity: 1},{productID: productID2,quantity: 2}}},wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewOrder(userID, tt.args.totalAmount, tt.args.products, orderAt)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(got,tt.want,cmp.AllowUnexported(Order{},OrderProduct{}),cmpopts.IgnoreFields(Order{},"id","orderAt"))
			if diff != "" {
				t.Errorf("NewOrder() = %v, want %v, error is %s",got,tt.want,err)
			}
		})
	}
}

func TestReconstract(t *testing.T) {
	type args struct {
		id          string
		userID      string
		totalAmount int64
		products    []OrderProduct
		now         time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    *Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Reconstract(tt.args.id, tt.args.userID, tt.args.totalAmount, tt.args.products, tt.args.now)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reconstract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reconstract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newOrder(t *testing.T) {
	type args struct {
		id          string
		userID      string
		totalAmount int64
		products    []OrderProduct
		now         time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    *Order
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newOrder(tt.args.id, tt.args.userID, tt.args.totalAmount, tt.args.products, tt.args.now)
			if (err != nil) != tt.wantErr {
				t.Errorf("newOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_UserID(t *testing.T) {
	type fields struct {
		id          string
		userID      string
		totalAmount int64
		product     OrderProducts
		orderAt     time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Order{
				id:          tt.fields.id,
				userID:      tt.fields.userID,
				totalAmount: tt.fields.totalAmount,
				product:     tt.fields.product,
				orderAt:     tt.fields.orderAt,
			}
			if got := o.UserID(); got != tt.want {
				t.Errorf("Order.UserID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_ID(t *testing.T) {
	type fields struct {
		id          string
		userID      string
		totalAmount int64
		product     OrderProducts
		orderAt     time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Order{
				id:          tt.fields.id,
				userID:      tt.fields.userID,
				totalAmount: tt.fields.totalAmount,
				product:     tt.fields.product,
				orderAt:     tt.fields.orderAt,
			}
			if got := o.ID(); got != tt.want {
				t.Errorf("Order.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_Totalamount(t *testing.T) {
	type fields struct {
		id          string
		userID      string
		totalAmount int64
		product     OrderProducts
		orderAt     time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Order{
				id:          tt.fields.id,
				userID:      tt.fields.userID,
				totalAmount: tt.fields.totalAmount,
				product:     tt.fields.product,
				orderAt:     tt.fields.orderAt,
			}
			if got := o.Totalamount(); got != tt.want {
				t.Errorf("Order.Totalamount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_Products(t *testing.T) {
	type fields struct {
		id          string
		userID      string
		totalAmount int64
		product     OrderProducts
		orderAt     time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   []OrderProduct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Order{
				id:          tt.fields.id,
				userID:      tt.fields.userID,
				totalAmount: tt.fields.totalAmount,
				product:     tt.fields.product,
				orderAt:     tt.fields.orderAt,
			}
			if got := o.Products(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.Products() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_OrderAt(t *testing.T) {
	type fields struct {
		id          string
		userID      string
		totalAmount int64
		product     OrderProducts
		orderAt     time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Order{
				id:          tt.fields.id,
				userID:      tt.fields.userID,
				totalAmount: tt.fields.totalAmount,
				product:     tt.fields.product,
				orderAt:     tt.fields.orderAt,
			}
			if got := o.OrderAt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.OrderAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_ProductIDs(t *testing.T) {
	type fields struct {
		id          string
		userID      string
		totalAmount int64
		product     OrderProducts
		orderAt     time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Order{
				id:          tt.fields.id,
				userID:      tt.fields.userID,
				totalAmount: tt.fields.totalAmount,
				product:     tt.fields.product,
				orderAt:     tt.fields.orderAt,
			}
			if got := o.ProductIDs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Order.ProductIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderProducts_ProductIDs(t *testing.T) {
	tests := []struct {
		name string
		o    OrderProducts
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.ProductIDs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderProducts.ProductIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderProducts_TotalAmount(t *testing.T) {
	tests := []struct {
		name string
		o    OrderProducts
		want int64
	}{
		{name: "正常系",o: OrderProducts{{price: 100,quantity: 1},{price: 200,quantity: 2}},want: 500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.o.TotalAmount(); got != tt.want {
				t.Errorf("OrderProducts.TotalAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewOrderProduct(t *testing.T) {
	productID := strings.RemoveHyphen(uuid.NewString())
	price := int64(100)
	type args struct {
		productID string
		price     int64
		quantity  int
	}
	tests := []struct {
		name    string
		args    args
		want    *OrderProduct
		wantErr bool
	}{
		{name: "正常系",args: args{productID: productID,price: price,quantity: 1},want: &OrderProduct{productID: productID,price: price,quantity: 1},wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewOrderProduct(tt.args.productID, tt.args.price, tt.args.quantity)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOrderProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(got,tt.want,cmp.AllowUnexported(OrderProduct{}))
			if diff != "" {
				t.Errorf("NewOrder() = %v, want %v, error is %s", got,tt.want,err)
			}
		})
	}
}

func TestOrderProduct_ProductID(t *testing.T) {
	type fields struct {
		productID string
		price     int64
		quantity  int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OrderProduct{
				productID: tt.fields.productID,
				price:     tt.fields.price,
				quantity:  tt.fields.quantity,
			}
			if got := o.ProductID(); got != tt.want {
				t.Errorf("OrderProduct.ProductID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderProduct_Price(t *testing.T) {
	type fields struct {
		productID string
		price     int64
		quantity  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OrderProduct{
				productID: tt.fields.productID,
				price:     tt.fields.price,
				quantity:  tt.fields.quantity,
			}
			if got := o.Price(); got != tt.want {
				t.Errorf("OrderProduct.Price() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderProduct_Quantity(t *testing.T) {
	type fields struct {
		productID string
		price     int64
		quantity  int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OrderProduct{
				productID: tt.fields.productID,
				price:     tt.fields.price,
				quantity:  tt.fields.quantity,
			}
			if got := o.Quantity(); got != tt.want {
				t.Errorf("OrderProduct.Quantity() = %v, want %v", got, tt.want)
			}
		})
	}
}
