package order

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	orderApp "github.com/o-ga09/tutorial-ec-backend/app/application/order"
	"github.com/o-ga09/tutorial-ec-backend/app/server/middleware"
	"github.com/o-ga09/tutorial-ec-backend/pkg/validator"
)

type handler struct {
	saveOrderUsecase *orderApp.SaveOrderUseCase
}

func NewHandlre(saveOrderUsecase *orderApp.SaveOrderUseCase) handler {
	return handler{
		saveOrderUsecase: saveOrderUsecase,
	}
}

type Sample struct {
	name string `validate:"min=1,max=99"`
	id string `validate:"max=9"`
}

// PostOrder godoc
// @Summary 商品を注文する
// @Tags order
// @Accept json
// @Produce json
// @Param request body []PostOrderParams ture "注文商品"
// @Success 201 {string} Response 
// @Router /v1/order [post]
func(o *handler) PostOrders(c *gin.Context) {
	var params []*PostOrderParams

	err := c.ShouldBindJSON(&params)
	if err != nil {
		slog.Log(c, middleware.SeverityError, "error","err msg",err)
		c.JSON(http.StatusBadRequest, Response{Code: 403, Message: "Bad Request"})
		return
	}

	validate := validator.GetValidator()
	for _, param := range params {
		if err := validate.Struct(param); err != nil {
			slog.Log(c, middleware.SeverityError, "error","err msg",params)
			c.JSON(http.StatusBadRequest, Response{Code: 403, Message: "Bad Request"})
			return
		}
	}

	// 本来はsessionのuserIDを使う
	u, _ := c.Get("user_id")
	userID, _ := u.(string)

	dtos := make([]orderApp.SaveOrderUseCaseInputDto,0,len(params))

	for _, dto := range params {
		dtos = append(dtos,orderApp.SaveOrderUseCaseInputDto{
			ProductID: dto.ProductID,
			Quantity: dto.Quantity,
		})
	}

	id, err := o.saveOrderUsecase.Run(c,userID,dtos,time.Now())
	if err != nil {
		slog.Log(c, middleware.SeverityError, "error","err msg",err)
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "Internal Server Error"})
		return
	}
	slog.Log(c, middleware.SeverityError, "success")
	c.JSON(http.StatusCreated, Response{Code: 201, Message: fmt.Sprintf("id : %s",id)})
}