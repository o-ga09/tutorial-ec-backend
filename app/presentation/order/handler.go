package order

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	orderApp "github.com/o-ga09/tutorial-ec-backend/app/application/order"
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

func(o *handler) PostOrders(c *gin.Context) {
	var params []*PostOrderParams

	err := c.ShouldBindJSON(&params)
	if err != nil {
		log.Println("error")
		c.JSON(http.StatusBadRequest, Response{Code: 403, Message: "error"})
		return
	}

	validate := validator.GetValidator()
	if err := validate.Struct(&params); err != nil {
		log.Println("error")
		c.JSON(http.StatusBadRequest, Response{Code: 403, Message: "error"})
		return
	}

	// 本来はsessionのuserIDを使う
	userID := "dummy_user_id"
	dtos := make([]orderApp.SaveOrderUseCaseInputDto,0,len(params))

	for _, dto := range params {
		dtos = append(dtos,orderApp.SaveOrderUseCaseInputDto{
			ProductID: dto.ProductID,
			Quantity: dto.Quantity,
		})
	}

	id, err := o.saveOrderUsecase.Run(c,userID,dtos,time.Now())
	if err != nil {
		log.Println("error")
		c.JSON(http.StatusInternalServerError, Response{Code: 500, Message: "error"})
		return
	}

	c.JSON(http.StatusCreated, Response{Code: 201, Message: fmt.Sprintf("id : %s",id)})
}