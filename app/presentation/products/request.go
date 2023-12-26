package products

type PostRequestParm struct {
	OwnerID     string `json:"owner_id,omitempty"`    // 出品者ID
	Name        string `json:"name,omitempty"`        // 商品名
	Description string `json:"description,omitempty"` // 商品の説明
	Price       int64  `json:"price,omitempty"`       // 商品金額
	Stock       int    `json:"stock,omitempty"`       // 商品在庫
}
