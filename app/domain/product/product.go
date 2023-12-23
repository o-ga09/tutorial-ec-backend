package product

import (
	"unicode/utf8"

	"github.com/google/uuid"

	errDomain "github.com/o-ga09/tutorial-go-fr/app/domain/error"
	"github.com/o-ga09/tutorial-go-fr/pkg/strings"
)

const (
	// 名前の最⼤値/最⼩値
	nameLengthMin = 1
	nameLengthMax = 100
	// 説明の最⼤値/最⼩値
	descriptionLengthMin = 1
	descriptionLengthMax = 1000
)

type Product struct {
	id          string // 商品ID
	ownerID     string // 出品者ID
	name        string // 商品名
	description string // 商品の説明
	price       int64  // 商品金額
	stock       int    // 商品在庫
}

func newProduct(id string, ownerID string, name string, description string, price int64, stock int) (*Product, error) {

	// 名前のバリデーション
	if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
		return nil, errDomain.INVALID_PRODUCTNAME
	}

	// 商品説明のバリデーション
	if utf8.RuneCountInString(description) < descriptionLengthMin || utf8.RuneCountInString(description) > descriptionLengthMax {
		return nil, errDomain.INVALID_PRODUCTDESC
	}

	// 価格のバリデーション
	if price < 1 {
		return nil, errDomain.INVALID_PRODUCTPRICE
	}

	// 在庫数のバリデーション
	// 在庫はないけど 、商品は登録したい等あるため 、0は許容する
	if stock < 0 {
		return nil, errDomain.INVALID_PRODUCTSTOCK
	}

	return &Product{
		id:          id,
		ownerID:     ownerID,
		name:        name,
		description: description,
		price:       price,
		stock:       stock,
	}, nil
}

func Reconstruct(id string, ownerID string, name string, description string, price int64, stock int) (*Product, error) {
	return &Product{
		id:          id,
		ownerID:     ownerID,
		name:        name,
		description: description,
		price:       price,
		stock:       stock,
	}, nil
}

func NewProduct(ownerID string, name string, description string, price int64, stock int) (*Product, error) {
	id := strings.RemoveHyphen(uuid.NewString())
	return newProduct(id,ownerID, name, description, price, stock)
}

func(p *Product) Consume(quantity int) error {
	if quantity < 0 {
		return errDomain.INVALID_PRODUCTSTOCK
	}

	if p.stock - quantity < 0 {
		return errDomain.INVALID_LACK_STOCK
	}

	p.stock -= quantity
	return nil
}

func(p *Product) ID() string {return p.id}
func(p *Product) OwnerID() string {return p.ownerID}
func(p *Product) Name() string {return p.name}
func(p *Product) Description() string {return p.description}
func(p *Product) Price() int64 {return p.price}
func(p *Product) Stock() int {return p.stock}