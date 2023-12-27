package cart

type PostCartsParams struct {
	ProductID string `json:"product_id,omitempty"`
	Quantity  int    `json:"quantity,omitempty"`
}
