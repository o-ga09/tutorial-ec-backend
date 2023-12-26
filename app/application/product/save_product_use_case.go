package product

import (
	"context"

	productDomain "github.com/o-ga09/tutorial-ec-backend/app/domain/product"
)

type SaveProductUsecase struct {
	productRepo productDomain.ProductRepository
}

func NewSaveProductUsecase(productRepo productDomain.ProductRepository) *SaveProductUsecase {
	return &SaveProductUsecase{productRepo: productRepo}
}

type SaveProductUsecaseInputDto struct {
	OwnerID     string // 出品者ID
	Name        string // 商品名
	Description string // 商品の説明
	Price       int64  // 商品金額
	Stock       int    // 商品在庫
}

type SaveProductUsecaseOutputDto struct {
	Id          string // 商品ID
	OwnerID     string // 出品者ID
	Name        string // 商品名
	Description string // 商品の説明
	Price       int64  // 商品金額
	Stock       int    // 商品在庫
}

func(u *SaveProductUsecase) Run(ctx context.Context, input SaveProductUsecaseInputDto) (*SaveProductUsecaseOutputDto, error) {
	p, err := productDomain.NewProduct(input.OwnerID,input.Name,input.Description,input.Price,input.Stock)
	if err != nil {
		return nil, err
	}

	if err := u.productRepo.Save(ctx,p); err != nil {
		return nil, err
	}
	
	return &SaveProductUsecaseOutputDto{
		p.ID(),
		p.OwnerID(),
		p.Name(),
		p.Description(),
		p.Price(),
		p.Stock(),
	}, nil
}