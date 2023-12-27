package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	userApp "github.com/o-ga09/tutorial-ec-backend/app/application/user"
)

type handler struct {
	findUserUsecase *userApp.FindUserUsecase
	saveUserUsecase *userApp.SaveUserUsecase
}

func NewHandler(findUserUsecase *userApp.FindUserUsecase, saveUserUsecase *userApp.SaveUserUsecase) handler {
	return handler{
		findUserUsecase: findUserUsecase,
		saveUserUsecase: saveUserUsecase,
	}
}

func(u handler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	dto, err := u.findUserUsecase.Run(c,id)
	if err != nil {
		log.Println("error")
		return
	}

	res := getUserResponse{
		User: userResponseModel{
			ID: dto.ID,
			Email: dto.Email,
			PhoneNumber: dto.Phonenumber,
			LastName: dto.Lastname,
			FirstName: dto.Firstname,
			Address: dto.Address,
		},
	}

	c.JSON(http.StatusOK, res)
}