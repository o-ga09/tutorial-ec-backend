package product

import "context"

type FetchProductlistDto struct {
	ID string
	ownerID     string // 出品者ID
	name        string // 商品名
	description string // 商品の説明
	price       int64  // 商品金額
	stock       int    // 商品在庫
}

type ProductQueryService interface {
	FetchProductList(ctx context.Context)([]*FetchProductlistDto, error)
}