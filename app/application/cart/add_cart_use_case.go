package cart

import (
	"context"
	"log/slog"

	cartDomain "github.com/o-ga09/tutorial-ec-backend/app/domain/cart"
	productDomain "github.com/o-ga09/tutorial-ec-backend/app/domain/product"
	"github.com/o-ga09/tutorial-ec-backend/app/server/middleware"
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
		slog.Log(ctx, middleware.SeverityError, "failed to get current cart","err msg",err)
		return err
	}

	// 在庫情報を取得
	product, err := u.productRepo.FindByID(ctx,dto.ProductID)
	if err != nil {
		slog.Log(ctx, middleware.SeverityError, "failed to get current product","err msg",err)
		return err
	}

	if err := product.Consume(dto.Quantity); err != nil {
		slog.Log(ctx, middleware.SeverityError, "failed to process consume","err msg",err)
		return err
	}

	// カートの更新
	if err := u.updateCart(cart, dto); err != nil {
		slog.Log(ctx, middleware.SeverityError, "failed to update cart","err msg",err)
		return err
	}

	// カートの永続化
	if err := u.cartRepo.Save(ctx,cart); err != nil {
		slog.Log(ctx, middleware.SeverityError, "failed to save cart","err msg",err)
		return err
	}
	
	return nil
}

func(u *AddCartUsecase) updateCart(cart *cartDomain.Cart, dto AddCartUsecaseInputDto) error {
	// 商品が0の時はカートから削除、それ以外は追加・更新
	if dto.Quantity == 0 {
		if err := cart.RemoveProduct(dto.ProductID); err != nil {
			slog.Log(context.Background(), middleware.SeverityError, "failed to remove cart","err msg",err)
			return nil
		}
		return nil
	}
	// カートに商品を追加
	if err := cart.AddProduct(dto.ProductID,dto.Quantity); err != nil {
		slog.Log(context.Background(), middleware.SeverityError, "failed to add cart","err msg",err)
		return err
	}

	return nil
}