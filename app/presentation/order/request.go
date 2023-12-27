package order

type PostOrderParams struct {
	ProductID string `json:"product_id,omitempty"`
	Quantity  int    `json:"quantity,omitempty"`
}
