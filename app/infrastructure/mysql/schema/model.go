package schema

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID      string    `gorm:"column:user_id;primary_key; type:varchar(255)" json:"user_id,omitempty"`
	TotalAmount int64     `gorm:"column:total_amount; type:varchar(255)" json:"total_amount,omitempty"`
	OrderAt     time.Time `gorm:"column:order_at; type:varchar(255)" json:"order_at,omitempty"`
}

type OrderProduct struct {
	gorm.Model
	OrderID   string `gorm:"column:order_id;primary_key; type:varchar(255)" json:"order_id,omitempty"`
	ProductID string `gorm:"column:product_id; type:varchar(255)" json:"product_id,omitempty"`
	Price     int64  `gorm:"column:price; type:varchar(255)" json:"price,omitempty"`
	Quantity  int    `gorm:"column:quantity; type:varchar(255)" json:"quantity,omitempty"`
}

type Product struct {
	gorm.Model
	ProductId   string `gorm:"column:product_id;primary_key; type:varchar(255)" json:"product_id"`   // 商品ID
	OwnerID     string `gorm:"column:owner_id; type:varchar(255)" json:"owner_id"`       // 出品者ID
	Name        string `gorm:"column:name; type:varchar(255)" json:"name"`               // 商品名
	Description string `gorm:"column:description; type:varchar(255)" json:"description"` // 商品の説明
	Price       int64  `gorm:"column:price; type:varchar(255)" json:"price"`             // 商品金額
	Stock       int    `gorm:"column:stock; type:varchar(255)" json:"stock"`             // 商品在庫
}

type User struct {
	gorm.Model
	UserID      string `gorm:"column:user_id;primary_key; type:varchar(255)" json:"user_id,omitempty"`
	Email       string `gorm:"column:email; type:varchar(255)" json:"email,omitempty"`
	PhoneNumber string `gorm:"column:phonenumber; type:varchar(255)" json:"phone_number,omitempty"`
	LastName    string `gorm:"column:last_name; type:varchar(255)" json:"last_name,omitempty"`
	FirstName   string `gorm:"column:first_name; type:varchar(255)" json:"first_name,omitempty"`
	Pref        string `gorm:"column:prefecture; type:varchar(255)" json:"pref,omitempty"`
	City        string `gorm:"column:city; type:varchar(255)" json:"city,omitempty"`
	Extra       string `gorm:"column:extra; type:varchar(255)" json:"extra,omitempty"`
}

type Owner struct {
	gorm.Model
	OwnerID string `gorm:"column:owner_id;primary_key; type:varchar(255)" json:"owner_id,omitempty"`
	Email   string `gorm:"column:email; type:varchar(255)" json:"email,omitempty"`
	Name    string `gorm:"column:name; type:varchar(255)" json:"name,omitempty"`
}
