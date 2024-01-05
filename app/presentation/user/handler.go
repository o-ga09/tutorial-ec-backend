package user

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	userApp "github.com/o-ga09/tutorial-ec-backend/app/application/user"
	"github.com/o-ga09/tutorial-ec-backend/app/server/middleware"
	"github.com/o-ga09/tutorial-ec-backend/pkg/validator"
	"gorm.io/gorm"
)

type handler struct {
	findUserUsecase *userApp.FindUserUsecase
	findAllUsersUsecase *userApp.FindAllUsersUseCase
	saveUserUsecase *userApp.SaveUserUsecase
	deleteUserUsercase *userApp.DeleteUserUsecase
}

func NewHandler(findUserUsecase *userApp.FindUserUsecase,findAllUsersUsecase *userApp.FindAllUsersUseCase, saveUserUsecase *userApp.SaveUserUsecase, deleteUserUsecase *userApp.DeleteUserUsecase) handler {
	return handler{
		findUserUsecase: findUserUsecase,
		findAllUsersUsecase: findAllUsersUsecase,
		saveUserUsecase: saveUserUsecase,
		deleteUserUsercase: deleteUserUsecase,
	}
}

// GetUser godoc
// @Summary ユーザーを取得する
// @Tags users
// @Accept json
// @Produce json
// @Param request path string ture "ユーザーID"
// @Success 200 {object} getUserResponse
// @Router /v1/products/:id [get]
func(u handler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	dto, err := u.findUserUsecase.Run(c,id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Log(c, middleware.SeverityInfo, "Not Found")
		c.JSON(http.StatusNotFound,gin.H{"code": 404, "message": "Not Found"})
		return
	}

	if err != nil {
		slog.Log(c, middleware.SeverityError, "error","err msg",err)
		c.JSON(http.StatusBadRequest,gin.H{"code": 500, "message": "Internal Server Error"})
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

func(h *handler) GetUsers(c *gin.Context) {
	dtos, err := h.findAllUsersUsecase.Run(c)
	if err != nil {
		slog.Log(c, middleware.SeverityError, "error","err msg",err)
		c.JSON(http.StatusBadRequest,gin.H{"code": 500, "message": "Internal Server Error"})
		return
	}

	res := []*getUserResponse{}
	for _, dto := range dtos {
		u := getUserResponse{
			User: userResponseModel{
				ID: dto.ID,
				Email: dto.Email,
				PhoneNumber: dto.Phonenumber,
				LastName: dto.Lastname,
				FirstName: dto.Firstname,
				Address: dto.Address,
			},
		}
		res = append(res, &u)
	}

	slog.Log(c, middleware.SeverityInfo,"success")
	c.JSON(http.StatusOK,res)
	return
}

// GetUser godoc
// @Summary ユーザーを登録する
// @Tags users
// @Accept json
// @Produce json
// @Param request body PostUserParam ture "ユーザー情報"
// @Success 201 {object} getUserResponse
// @Router /v1/products [post]
func(h *handler) CreateUser(c *gin.Context) {
	var param *PostUserParam

	err := c.ShouldBindJSON(&param)
	if err != nil {
		slog.Log(c, middleware.SeverityError, "param bind error","err msg",err)
		c.JSON(http.StatusBadRequest,gin.H{"code": 403, "message": "Bad Request"})
		return
	}

	validate := validator.GetValidator()
	if validate.Struct(&param); err != nil {
		slog.Log(c, middleware.SeverityError, "param validation error","err msg",err)
		c.JSON(http.StatusBadRequest,gin.H{"code": 403, "message": "Bad Request"})
		return
	}

	user := userApp.SaveUserUsecaseDto{
		Email: param.Email,
		Phonenumber: param.PhoneNumber,
		Lastname: param.LastName,
		Firstname: param.FirstName,
		Pref: param.Pref,
		City: param.City,
		Extra: param.Extra,
	}

	if err := h.saveUserUsecase.Run(c,user); err != nil {
		slog.Log(c, middleware.SeverityError, "failed to create user","err msg",err)
		c.JSON(http.StatusInternalServerError,gin.H{"code": 500, "message": "Internal Server Error"})
	}

	slog.Log(c, middleware.SeverityInfo, "success")
	c.JSON(http.StatusOK,gin.H{"code": 201, "message": "user created"})
	return
}

func(h *handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := h.deleteUserUsercase.Run(c,id)
	if err != nil {
		slog.Log(c, middleware.SeverityError, "failed to delete")
		c.JSON(http.StatusNotFound,gin.H{"code": 500, "message": "Internal Server Error"})
		return
	}

	slog.Log(c, middleware.SeverityInfo, "success")
	c.JSON(http.StatusOK,gin.H{"code": 201, "message": "delete"})
	return
}