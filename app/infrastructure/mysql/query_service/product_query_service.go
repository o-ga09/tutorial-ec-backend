package queryservice

import (
	"context"

	"github.com/o-ga09/tutorial-ec-backend/app/application/product"
	"gorm.io/gorm"
)

type productQueryService struct{
	conn *gorm.DB
}

type FetchProductlist struct {
    ID          string
    ownerID     string // 出品者ID
    name        string // 商品名
    description string // 商品の説明
    price       int64  // 商品金額
    stock       int    // 商品在庫
}

// FetchProductList implements product.ProductQueryService.
func (qs *productQueryService) FetchProductList(ctx context.Context) ([]*product.FetchProductlistDto, error) {
	res := []*FetchProductlist{}
	err := qs.conn.Find(&res).Error
	if err != nil {
		return nil, err
	}

	var productFetchSerbiceDto []*product.FetchProductlistDto
	for _, r := range res {
		productFetchSerbiceDto = append(productFetchSerbiceDto, &product.FetchProductlistDto{
			ID: r.ID,
			OwnerID: r.ownerID,
			Name: r.name,
			Description: r.description,
			Price: r.price,
			Stock: r.stock,
		})
	}
	return productFetchSerbiceDto, nil
}

func NewproductQueryService(conn *gorm.DB) product.ProductQueryService {
	return &productQueryService{
		conn: conn,
	}
}
