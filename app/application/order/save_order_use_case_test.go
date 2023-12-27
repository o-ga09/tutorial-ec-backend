package order

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	cartDomain "github.com/o-ga09/tutorial-ec-backend/app/domain/cart"
	orderDomain "github.com/o-ga09/tutorial-ec-backend/app/domain/order"
	"github.com/o-ga09/tutorial-ec-backend/pkg/strings"
)

func TestNewSaveOrderUseCase(t *testing.T) {
	MockOrderDomainService := orderDomain.OrderDomainServiceMock{}
	MockCartRepository := cartDomain.CartRepositoryMock{}
	type args struct {
		ods orderDomain.OrderDomainService
		cr  cartDomain.CartRepository
	}
	tests := []struct {
		name string
		args args
		want *SaveOrderUseCase
	}{
		{name: "正常系",args: args{ods: &MockOrderDomainService,cr: &MockCartRepository},want: &SaveOrderUseCase{&MockOrderDomainService,&MockCartRepository}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSaveOrderUseCase(tt.args.ods, tt.args.cr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSaveOrderUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveOrderUseCase_Run(t *testing.T) {
	orderID := strings.RemoveHyphen(uuid.NewString())
	userID := strings.RemoveHyphen(uuid.New().String())
	cart, _ := cartDomain.NewCart(userID)
	
	MockOrderDomainService := orderDomain.OrderDomainServiceMock{
		OrderProductsFunc: func(ctx context.Context, cart *cartDomain.Cart, now time.Time) (string, error) {
			return orderID, nil
		},
	}
	MockCartRepository := cartDomain.CartRepositoryMock{
		FindByIDFunc: func(ctx context.Context, userID string) (*cartDomain.Cart, error) {
			return cart, nil
		},
	}
	usecase := NewSaveOrderUseCase(&MockOrderDomainService,&MockCartRepository)
	now := time.Now()
	dtos := []SaveOrderUseCaseInputDto{
		{ProductID: strings.RemoveHyphen(uuid.NewString()),Quantity: 1},
		{ProductID: strings.RemoveHyphen(uuid.NewString()),Quantity: 3},
	}
	for _, dto := range dtos {
		cart.AddProduct(dto.ProductID,dto.Quantity)
	}

	tests := []struct {
		name    string
		dtos    []SaveOrderUseCaseInputDto
		wantErr bool
	}{
		{
			name: "正常系",
			dtos: dtos,
			wantErr: false,
		},
			
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := usecase.Run(context.Background(),userID,dtos,now)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveOrderUseCase.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != orderID {
				t.Errorf("SaveOrderUseCase.Run() = %v, want %v", got, orderID)
			}
		})
	}
}

func TestSaveOrderUseCase_getValidCart(t *testing.T) {
	userID := strings.RemoveHyphen(uuid.New().String())
	cart, _ := cartDomain.NewCart(userID)
	dtos := []SaveOrderUseCaseInputDto{
		{ProductID: strings.RemoveHyphen(uuid.NewString()),Quantity: 1},
		{ProductID: strings.RemoveHyphen(uuid.NewString()),Quantity: 3},
	}
	for _, dto := range dtos {
		cart.AddProduct(dto.ProductID,dto.Quantity)
	}

	MockOrderDomainService := orderDomain.OrderDomainServiceMock{}
	MockCartRepository := cartDomain.CartRepositoryMock{
		FindByIDFunc: func(ctx context.Context, userID string) (*cartDomain.Cart, error) {
			return cart, nil
		},
	}
	usecase := NewSaveOrderUseCase(&MockOrderDomainService,&MockCartRepository)

	tests := []struct {
		name    string
		want    *cartDomain.Cart
		wantErr bool
	}{
		{name: "正常系",want: cart,wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := usecase.getValidCart(context.Background(),userID,dtos)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveOrderUseCase.getValidCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SaveOrderUseCase.getValidCart() = %v, want %v", got, tt.want)
			}
		})
	}
}
