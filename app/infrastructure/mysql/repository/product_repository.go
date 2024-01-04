package repository

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/o-ga09/tutorial-ec-backend/app/domain/product"
	"github.com/o-ga09/tutorial-ec-backend/app/infrastructure/mysql/schema"
	"github.com/o-ga09/tutorial-ec-backend/app/server/middleware"
	"gorm.io/gorm"
)

type productRepository struct{
	conn *gorm.DB
}

// FindByID implements product.ProductRepository.
func (r *productRepository) FindByID(ctx context.Context, id string) (*product.Product, error) {

	repoProduct := schema.Product{}
	if err := r.conn.Where("`id` = ?", id).First(&repoProduct).Error; err != nil {
		return nil, err
	}

	fmt.Println(repoProduct.OwnerID)
	product, err := product.Reconstruct(repoProduct.ProductId,repoProduct.OwnerID,repoProduct.Name,repoProduct.Description,repoProduct.Price,repoProduct.Stock)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// FindByIDs implements product.ProductRepository.
func (r *productRepository) FindByIDs(ctx context.Context, ids []string) ([]product.Product, error) {
	res := []product.Product{}
	repoProducts := []schema.Product{}

	if err := r.conn.Where("id IN (?)",ids).Find(&repoProducts).Error; err != nil {
		return nil, err
	}

	for _, repoProduct := range repoProducts {
		product, err := product.Reconstruct(repoProduct.ProductId,repoProduct.OwnerID,repoProduct.Name,repoProduct.Description,repoProduct.Price,repoProduct.Stock)
		if err != nil {
			return nil, err
		}
		res = append(res, *product)
	}

	return res, nil
}

// Save implements product.ProductRepository.
func (r *productRepository) Save(ctx context.Context, product *product.Product) error {
	repoProduct := schema.Product{
		ProductId: product.ID(),
		OwnerID: product.OwnerID(),
		Name: product.Name(),
		Description: product.Description(),
		Price: product.Price(),
		Stock: product.Stock(),
	}

	err := r.conn.Create(&repoProduct).Error
	if err != nil {
		slog.Log(ctx,middleware.SeverityError,"repository error",err)
		return err
	}
	return nil
}

func (r *productRepository) Update(ctx context.Context, product *product.Product) error {
	repoProduct := schema.Product{
		ProductId: product.ID(),
		OwnerID: product.OwnerID(),
		Name: product.Name(),
		Description: product.Description(),
		Price: product.Price(),
		Stock: product.Stock(),
	}

	r.conn.Save(repoProduct)
	return nil
}

func (r *productRepository) Delete(ctx context.Context, product *product.Product) error {
	r.conn.Where("id = ?",product.ID()).Delete(schema.Product{})
	return nil
}

func NewProductRepository(conn *gorm.DB) product.ProductRepository {
	return &productRepository{conn: conn}
}
