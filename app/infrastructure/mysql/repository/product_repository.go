package repository

import (
	"context"
	"fmt"

	"github.com/o-ga09/tutorial-ec-backend/app/domain/product"
	"gorm.io/gorm"
)

type productRepository struct{
	conn *gorm.DB
}
type Product struct {
	Id          string `gorm:"id,omitempty"`          // 商品ID
	OwnerID     string `gorm:"owner_id,omitempty"`    // 出品者ID
	Name        string `gorm:"name,omitempty"`        // 商品名
	Description string `gorm:"description,omitempty"` // 商品の説明
	Price       int64  `gorm:"price,omitempty"`       // 商品金額
	Stock       int    `gorm:"stock,omitempty"`       // 商品在庫
}
// FindByID implements product.ProductRepository.
func (r *productRepository) FindByID(ctx context.Context, id string) (*product.Product, error) {

	repoProduct := Product{}
	if err := r.conn.Where("`id` = ?", id).First(&repoProduct).Error; err != nil {
		return nil, err
	}

	fmt.Println(repoProduct.OwnerID)
	product, err := product.Reconstruct(repoProduct.Id,repoProduct.OwnerID,repoProduct.Name,repoProduct.Description,repoProduct.Price,repoProduct.Stock)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// FindByIDs implements product.ProductRepository.
func (r *productRepository) FindByIDs(ctx context.Context, ids []string) ([]product.Product, error) {
	res := []product.Product{}
	repoProducts := []Product{}

	if err := r.conn.Where("id IN (?)",ids).Find(&repoProducts).Error; err != nil {
		return nil, err
	}

	for _, repoProduct := range repoProducts {
		product, err := product.Reconstruct(repoProduct.Id,repoProduct.OwnerID,repoProduct.Name,repoProduct.Description,repoProduct.Price,repoProduct.Stock)
		if err != nil {
			return nil, err
		}
		res = append(res, *product)
	}

	return res, nil
}

// Save implements product.ProductRepository.
func (r *productRepository) Save(ctx context.Context, product *product.Product) error {
	repoProduct := Product{
		Id: product.ID(),
		OwnerID: product.OwnerID(),
		Name: product.Name(),
		Description: product.Description(),
		Price: product.Price(),
		Stock: product.Stock(),
	}

	r.conn.Create(repoProduct)
	return nil
}

func (r *productRepository) Update(ctx context.Context, product *product.Product) error {
	repoProduct := Product{
		Id: product.ID(),
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
	r.conn.Where("id = ?",product.ID()).Delete(Product{})
	return nil
}

func NewProductRepository(conn *gorm.DB) product.ProductRepository {
	return &productRepository{conn: conn}
}
