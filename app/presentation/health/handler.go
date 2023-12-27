package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {}

func NewHandler() *handler {
	return &handler{}
}

// HealthCheck godoc
// @Summary ヘルスチェック
// @Tags healthcheck
// @Accept json
// @Produce json
// @Success 200 {object} Response
// @Router /v1/health [get]
func(h *handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK,Response{Status: "OK"})
	return
}