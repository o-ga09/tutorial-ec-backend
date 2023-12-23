package order

import (
	"context"
	"time"

	cartDomain "github.com/o-ga09/tutorial-go-fr/app/domain/cart"
	errDomain "github.com/o-ga09/tutorial-go-fr/app/domain/error"
	productDoamin "github.com/o-ga09/tutorial-go-fr/app/domain/product"
)

type orderDomainService struct {
	orderRepo   OrderRepository
	productRepo productDoamin.ProductRepository
}


func NewOrderDomainService(orderRepo OrderRepository, productRepo productDoamin.ProductRepository) OrderDomainService {
	return &orderDomainService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (s *orderDomainService)OrderProducts(ctx context.Context, cart *cartDomain.Cart, now time.Time) (string, error) {
	// 注文対象の商品を取得
	ps, err := s.productRepo.FindByIDs(ctx,cart.ProductIDs())
	if err != nil {
		return "", err
	}

	productMap := make(map[string]*productDoamin.Product)
	for _, p := range ps {
		productMap[p.ID()] = &p
	}

	// 注文処理
	ops := make([]OrderProduct, 0,len(cart.ProductIDs()))
	for _, cp := range cart.Products() {
		p, ok := productMap[cp.ProductID()]
		op, err := NewOrderProduct(cp.ProductID(),p.Price(),cp.Quantity())
		if err != nil {
			return "" , err
		}

		ops = append(ops, *op)
		if !ok {
			return "", errDomain.NOT_FOUND
		}

		if err := p.Consume(cp.Quantity()); err != nil {
			return "", err
		}

		if err := s.productRepo.Save(ctx, p); err != nil {
			return "", err
		}
	}

	// 注文履歴保存
	o, err := NewOrder(cart.UserID(), OrderProducts(ops).TotalAmount(), ops, now)
	if err != nil {
		return "" , err
	}

	if err := s.orderRepo.Save(ctx, o); err != nil {
 		return "", err
	}

	return o.ID(), nil
}