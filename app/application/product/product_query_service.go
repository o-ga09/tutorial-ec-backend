package product

import "context"

type FetchProductlistDto struct {
	ID string
	OwnerID     string // 出品者ID
	Name        string // 商品名
	Description string // 商品の説明
	Price       int64  // 商品金額
	Stock       int    // 商品在庫
}

type ProductQueryService interface {
	FetchProductList(ctx context.Context)([]*FetchProductlistDto, error)
}