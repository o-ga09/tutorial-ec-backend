package repository

import (
	"context"
	"encoding/json"
	"fmt"

	cartDomain "github.com/o-ga09/tutorial-ec-backend/app/domain/cart"
	"github.com/redis/go-redis/v9"
)

type cartRepository struct{
	client *redis.Client
}

type cartProduct struct {
	ProductID string `json:"product_id,omitempty"`
	Quantity  int    `json:"quantity,omitempty"`
}

// FindByID implements cart.CartRepository.
func (c *cartRepository) FindByID(ctx context.Context, userID string) (*cartDomain.Cart, error) {
	cart, err := cartDomain.NewCart(userID)
	if err != nil {
		return nil, err
	}

	// userIDをキーにしたカート情報がエラー
	jsonData, err := c.client.Get(ctx, cartKey(userID)).Result()
	if err != nil {
		if err == redis.Nil {
			// キーがなかった場合は空のカートを返す
			return cart, nil
		}
		return nil, err
	}

	// 取得した JSON データを CartProduct のスライスにデシリアル化
	var products []cartProduct
	err = json.Unmarshal([]byte(jsonData), &products)
	if err != nil {
		return nil, err
	}

	// Redisから取得したカート情報をCartドメインに変換
	for _, p := range products {
		cart.AddProduct(p.ProductID, p.Quantity)
	}
	return cart, nil
}

// Save implements cart.CartRepository.
func (c *cartRepository) Save(ctx context.Context, cart *cartDomain.Cart) error {
	// カート情報をRedisに保存
	cps := make([]*cartProduct, 0, len(cart.Products()))
	for _, cp := range cart.Products() {
		cps = append(cps, &cartProduct{
			ProductID: cp.ProductID(),
			Quantity:  cp.Quantity(),
		})
	}
	j, err := json.Marshal(cps)
	if err != nil {
		return err
	}
	key := cartKey(cart.UserID())
	if _, err := c.client.Set(ctx, key, j, cartDomain.CART_TIME_OUT).Result(); err != nil {
		return err
	}
	return nil
}

func NewCartRepository(client *redis.Client) cartDomain.CartRepository {
	return &cartRepository{
		client: client,
	}
}

func cartKey(userId string) string {
	return fmt.Sprintf("cart-UserID-%s",userId)
}