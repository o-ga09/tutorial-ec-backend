package cart

import (
	"time"

	errDomain "github.com/o-ga09/tutorial-go-fr/app/domain/error"
)

// ================================ //
// ユーザーが商品を一時的に保存するカート //
// ユーザー毎に1つ存在する             //
// カートの中には、複数の商品が入っている //
// ================================ //

const CART_TIME_OUT = time.Minute * 30

type cartProduct struct {
	productID string
	quantity int
}

type Cart struct {
	userID string
	products []cartProduct
}

func(c *cartProduct) ProductID() string {
	return c.productID
}

func(c *cartProduct) Quantity() int {
	return c.quantity
}

func newCartProduct(productID string , quantity int) (*cartProduct, error) {
	
	
	return &cartProduct{productID: productID, quantity: quantity}, nil
}

func NewCart(userID string) (*Cart ,error) {
	return &Cart{userID: userID, products: []cartProduct{}}, nil
}

func(c *Cart) UserID() string {
	return c.userID
}

func(c *Cart) Products() []cartProduct {
	return c.products
}

func(c *Cart) ProductIDs() []string {
	var productIDs []string
	for _, p := range c.products {
		productIDs = append(productIDs,p.productID)
	}
	return productIDs
}

func(c *Cart) QuantityByProductID(productID string) (int, error) {
	for _, p := range c.products {
		if p.productID == productID {
			return p.quantity, nil
		}
	}
	return 0, errDomain.NOT_FOUND
}

func(c *Cart) AddProduct(productID string, quantity int) error {
	cp, err := newCartProduct(productID, quantity)
	if err != nil {
		return err
	}

	for i, p := range c.products {
		if p.productID == productID {
			c.products[i] = *cp
			return nil
		}
	}

	c.products= append(c.products, *cp)
	return nil
}

func(c *Cart) RemoveProduct(productID string) error {
	var newProduct []cartProduct
	for _, p := range c.products {
		if p.productID == productID {
			continue
		}
		newProduct = append(newProduct, p)
	}

	c.products = newProduct
	return nil
}