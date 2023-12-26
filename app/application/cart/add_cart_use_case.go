package cart

import (
	"context"

	cartDomain "github.com/o-ga09/tutorial-ec-backend/app/domain/cart"
	productDomain "github.com/o-ga09/tutorial-ec-backend/app/domain/product"
)

type AddCartUsecase struct {
	cartRepo cartDomain.CartRepository
	productRepo productDomain.ProductRepository
}

func NewAddCartUsecase(cartRepo cartDomain.CartRepository, productRepo productDomain.ProductRepository) *AddCartUsecase {
	return &AddCartUsecase{
		cartRepo: cartRepo,
		productRepo: productRepo,
	}
}

type AddCartUsecaseInputDto struct {
	UserID string
	ProductID string
	Quantity int
}

func(u *AddCartUsecase) Run(ctx context.Context, dto AddCartUsecaseInputDto) error {
	// 現在のカート一覧を取得
	cart, err := u.cartRepo.FindByID(ctx,dto.UserID)
	if err != nil {
		return err
	}

	// 在庫情報を取得
	product, err := u.productRepo.FindByID(ctx,dto.ProductID)
	if err != nil {
		return err
	}

	if err := product.Consume(dto.Quantity); err != nil {
		return err
	}

	// カートの更新
	if err := u.updateCart(cart, dto); err != nil {
		return err
	}

	// カートの永続化
	if err := u.cartRepo.Save(ctx,cart); err != nil {
		return err
	}
	
	return nil
}

func(u *AddCartUsecase) updateCart(cart *cartDomain.Cart, dto AddCartUsecaseInputDto) error {
	// 商品が0の時はカートから削除、それ以外は追加・更新
	if dto.Quantity == 0 {
		if err := cart.RemoveProduct(dto.ProductID); err != nil {
			return nil
		}
		return nil
	}
	// カートに商品を追加
	if err := cart.AddProduct(dto.ProductID,dto.Quantity); err != nil {
		return err
	}

	return nil
}