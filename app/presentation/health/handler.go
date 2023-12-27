package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type handler struct {}

func NewHandler() *handler {
	return &handler{}
}

func(h *handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK,Response{Status: "OK"})
	return
}