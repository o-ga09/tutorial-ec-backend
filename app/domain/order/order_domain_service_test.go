package order

import (
	"context"
	"reflect"
	"testing"
	"time"

	cartDomain "github.com/o-ga09/tutorial-ec-backend/app/domain/cart"
	productDoamin "github.com/o-ga09/tutorial-ec-backend/app/domain/product"
)

func TestNewOrderDomainService(t *testing.T) {
	type args struct {
		orderRepo   OrderRepository
		productRepo productDoamin.ProductRepository
	}
	tests := []struct {
		name string
		args args
		want OrderDomainService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrderDomainService(tt.args.orderRepo, tt.args.productRepo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrderDomainService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_orderDomainService_OrderProducts(t *testing.T) {
	type fields struct {
		orderRepo   OrderRepository
		productRepo productDoamin.ProductRepository
	}
	type args struct {
		ctx  context.Context
		cart *cartDomain.Cart
		now  time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &orderDomainService{
				orderRepo:   tt.fields.orderRepo,
				productRepo: tt.fields.productRepo,
			}
			got, err := s.OrderProducts(tt.args.ctx, tt.args.cart, tt.args.now)
			if (err != nil) != tt.wantErr {
				t.Errorf("orderDomainService.OrderProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("orderDomainService.OrderProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}
