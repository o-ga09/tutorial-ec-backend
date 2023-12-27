package cart

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	cartApp "github.com/o-ga09/tutorial-ec-backend/app/application/cart"
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

func(h handler) PostCart(c *gin.Context) {
	var params PostCartsParams

	if err := c.ShouldBindJSON(&params); err != nil {
		log.Println("error")
		c.JSON(http.StatusBadRequest, Response{Code: 403, Message: "error"})
	}

	validate := validator.GetValidator()
	if err := validate.Struct(&params); err != nil {
		log.Println("error")
		c.JSON(http.StatusBadRequest, Response{Code: 403, Message: "error"})
	}

	// 本来はsessionのuserIDを使う
	userID := "summy_user_id"

	dto := cartApp.AddCartUsecaseInputDto{
		ProductID: params.ProductID,
		Quantity: params.Quantity,
		UserID: userID,
	}

	if err := h.addCartUsecase.Run(c,dto); err != nil {
		log.Println("error")
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "error"})
	}

	c.JSON(http.StatusNoContent, Response{Code: 204, Message: "OK"})
}