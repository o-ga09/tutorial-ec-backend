package cart

import (
	"testing"

	"github.com/google/uuid"
	"github.com/o-ga09/tutorial-go-fr/pkg/strings"
)

func TestCart_QuantityByProductID(t *testing.T) {
	userID := strings.RemoveHyphen(uuid.New().String())
	productID1 := strings.RemoveHyphen(uuid.New().String())
	productID2 := strings.RemoveHyphen(uuid.New().String())

	cart := &Cart{
		userID: userID,
		products: []cartProduct{
			{
				productID: productID1,
				quantity: 1,
			},
			{
				productID: productID2,
				quantity: 2,
			},
		},
	}

	tests := []struct{
		name string
		targetProductID string
		want int
		wantErr bool
	}{
		{name: "正常系",targetProductID: productID1,want: 1,wantErr: false},
		{name: "カートの中に商品がない場合はエラーを返す",targetProductID: "test",want: 0,wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cart.QuantityByProductID(tt.targetProductID)
			if (err != nil ) != tt.wantErr {
				t.Errorf("cart.QuantityByProductID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("cart.QuantityByproductID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCart_AddProduct(t *testing.T) {
	userID := strings.RemoveHyphen(uuid.New().String())
	productID1 := strings.RemoveHyphen(uuid.New().String())
	productID2 := strings.RemoveHyphen(uuid.New().String())
	newproductID := strings.RemoveHyphen(uuid.NewString())

	cart := &Cart{
		userID: userID,
		products: []cartProduct{
			{
				productID: productID1,
				quantity: 1,
			},
			{
				productID: productID2,
				quantity: 2,
			},
		},
	}

	type args struct{
		productID string
		quantity int
	}

	tests := []struct{
		name string
		args args
		want *Cart
		wantErr bool
	}{
		{name: "正常系 - カートに存在しない商品を追加する ",args: args{productID: newproductID, quantity: 1},want: &Cart{userID: userID, products: []cartProduct{{productID: productID1,quantity: 1},{productID: productID2,quantity: 2},{productID: newproductID,quantity: 1}}},wantErr: false},
		{name: "正常系 - 商品がすでにカートに入っている場合は更新する",args: args{productID: productID1, quantity: 3},want: &Cart{userID: userID, products: []cartProduct{{productID: productID1,quantity: 4},{productID: productID2,quantity: 2},{productID: newproductID,quantity: 1}}},wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := cart.AddProduct(tt.args.productID,tt.args.quantity); (err != nil) != tt.wantErr {
				t.Errorf("cart.AddProduct() error %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCart_RemoveProduct(t *testing.T) {
	userID := strings.RemoveHyphen(uuid.New().String())
	productID1 := strings.RemoveHyphen(uuid.New().String())
	productID2 := strings.RemoveHyphen(uuid.New().String())

	cart := &Cart{
		userID: userID,
		products: []cartProduct{
			{
				productID: productID1,
				quantity: 1,
			},
			{
				productID: productID2,
				quantity: 2,
			},
		},
	}

	tests := []struct{
		name string
		targetProductID string
		want *Cart
		wantErr bool
	}{
		{name: "正常系",targetProductID: productID1,want: &Cart{userID: userID,products: []cartProduct{{productID: productID2, quantity: 2}}},wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := cart.RemoveProduct(tt.targetProductID); (err != nil) != tt.wantErr {
				t.Errorf("cart.RemoveProduct() error %v, wantErr %v",err,tt.wantErr)
			}
		})
	}
}