package order

import (
	"context"
	"time"

	cartDomain "github.com/o-ga09/tutorial-go-fr/app/domain/cart"
	errDomain "github.com/o-ga09/tutorial-go-fr/app/domain/error"
	orderDomain "github.com/o-ga09/tutorial-go-fr/app/domain/order"
)

type SaveOrderUseCase struct {
	orderDomainService orderDomain.OrderDomainService
	cartRepo cartDomain.CartRepository
}

type SaveOrderUseCaseInputDto struct {
	productID string
	Quantity int
}

func NewSaveOrderUseCase(ods orderDomain.OrderDomainService, cr cartDomain.CartRepository) *SaveOrderUseCase {
	return &SaveOrderUseCase{
		orderDomainService: ods,
		cartRepo: cr,
	}
}

func(s *SaveOrderUseCase) Run(ctx context.Context, userID string, dtos []SaveOrderUseCaseInputDto, now time.Time) (string, error) {
	// カートから商品情報を取得
	cart, err := s.getValidCart(ctx, userID, dtos)
	if err != nil {
		return "", err
	}

	// 注文処理
	orderID, err := s.orderDomainService.OrderProducts(ctx,cart,now)
	if err != nil {
		return "", err
	}

	return orderID, nil
}

func(s *SaveOrderUseCase) getValidCart(ctx context.Context,userID string, dtos []SaveOrderUseCaseInputDto) (*cartDomain.Cart, error) {
	// カートから商品情報を取得
	cart, err := s.cartRepo.FindByID(ctx,userID)
	if err != nil {
		return nil, err
	}

	for _, dto := range dtos {
		pq, err := cart.QuantityByProductID(dto.productID)
		if err != nil {
			return nil, err
		}
		// DTOで渡ってきた数量と、カートの数量が一致しない場合はエラ＝
		if pq != dto.Quantity {
			return nil, errDomain.INVALID_QUANTITY
		}
	}
	
	return cart, nil
}