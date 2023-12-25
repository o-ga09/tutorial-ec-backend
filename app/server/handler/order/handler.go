package order

import "gofr.dev/pkg/gofr"

type OrderHandler struct {}

func(h *OrderHandler) OrderHandler(c *gofr.Context) (interface{}, error) {
	c.Gofr.Logger.Log("OK")
	return "OK", nil
}