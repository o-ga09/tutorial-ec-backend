package error

import "errors"

var (
	INVALID_OWNERID = errors.New("オーナーIDの値が不正です 。")
	INVALID_PRODUCTNAME = errors.New("商品名の値が不正です")
	INVALID_PRODUCTDESC = errors.New("商品説明の値が不正です 。")
	INVALID_PRODUCTPRICE = errors.New("価格の値が不正です")
	INVALID_PRODUCTSTOCK = errors.New("在庫数の値が不正です 。")
	INVALID_LACK_STOCK = errors.New("在庫数が不足しています 。")

	NOT_FOUND = errors.New("商品が見つかりません。")
	INVALID_QUANTITY = errors.New("カートの商品数と一致しません。")
)

type Error struct {
	description string
}

func(e *Error) Error() string {
	return e.description
}

func NewError(s string) *Error {
	return &Error{
		description: s,
	}
}