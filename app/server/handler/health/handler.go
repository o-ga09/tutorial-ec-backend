package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {}

func(h *HealthHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{"Status": "OK"})
	return
}

func NewHealthCheckHandler() *HealthHandler {
	return &HealthHandler{}
}