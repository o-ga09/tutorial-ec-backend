package cart

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	cartApp "github.com/o-ga09/tutorial-ec-backend/app/application/cart"
	"github.com/o-ga09/tutorial-ec-backend/app/server/middleware"
	"github.com/o-ga09/tutorial-ec-backend/pkg/validator"
)

type handler struct {
	addCartUsecase *cartApp.AddCartUsecase
}

func NewHandler(addCartUsecase *cartApp.AddCartUsecase) handler {
	return handler{
		addCartUsecase: addCartUsecase,
	}
}

// PostCart godoc
// @Summary カートに商品を保存する
// @Tags cart
// @Accept json
// @Produce json
// @Param request body PostCartsParams ture "カートの商品"
// @Success 204 {object} Response
// @Router /v1/cart [post]
func(h handler) PostCart(c *gin.Context) {
	var params PostCartsParams

	if err := c.ShouldBindJSON(&params); err != nil {
		slog.Log(c, middleware.SeverityError, "error","err msg",err)
		c.JSON(http.StatusBadRequest, Response{Code: 403, Message: "error"})
	}

	validate := validator.GetValidator()
	if err := validate.Struct(&params); err != nil {
		slog.Log(c, middleware.SeverityError, "error","err msg",err)
		c.JSON(http.StatusBadRequest, Response{Code: 403, Message: "Bad Request"})
		return
	}

	u, _ := c.Get("user_id")
	userID, _ := u.(string)

	dto := cartApp.AddCartUsecaseInputDto{
		ProductID: params.ProductID,
		Quantity: params.Quantity,
		UserID: userID,
	}

	if err := h.addCartUsecase.Run(c,dto); err != nil {
		slog.Log(c, middleware.SeverityError, "error","err msg",err)
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "Internal Server Error"})
		return
	}

	c.JSON(http.StatusNoContent, Response{Code: 204, Message: "OK"})
}