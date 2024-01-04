package queryservice

import (
	"context"
	"log/slog"

	"github.com/o-ga09/tutorial-ec-backend/app/application/product"
	"github.com/o-ga09/tutorial-ec-backend/app/server/middleware"
	"gorm.io/gorm"
)

type productQueryService struct{
	conn *gorm.DB
}

type FetchProductlist struct {
    ProductID string `gorm:"column:product_id"`
	ProductName string `gorm:"column:name"`
	ProductDescription string `gorm:"column:description"`
	ProductPrice int64 `gorm:"column:price"`
	ProductStock int `gorm:"column:stock"`
	OwnerID string `gorm:"column:owner_id"`
	OwnerName string `gorm:"column:name"`
}

// FetchProductList implements product.ProductQueryService.
func (qs *productQueryService) FetchProductList(ctx context.Context) ([]*product.FetchProductlistDto, error) {
	res := []FetchProductlist{}
	err := qs.conn.Debug().Table("products").Select("products.*, owners.*").Joins("JOIN owners ON products.owner_id = owners.owner_id").Scan(&res).Error
	if err != nil {
		return nil, err
	}

	var productFetchSerbiceDto []*product.FetchProductlistDto
	for _, r := range res {
		productFetchSerbiceDto = append(productFetchSerbiceDto, &product.FetchProductlistDto{
			ID: r.ProductID,
			OwnerID: r.OwnerID,
			Name: r.ProductName,
			Description: r.ProductDescription,
			Price: r.ProductPrice,
			Stock: r.ProductStock,
			OwnerName: r.OwnerName,
		})
	}
	slog.Log(ctx, middleware.SeverityInfo, "done","reult",res)
	return productFetchSerbiceDto, nil
}

func NewproductQueryService(conn *gorm.DB) product.ProductQueryService {
	return &productQueryService{
		conn: conn,
	}
}
