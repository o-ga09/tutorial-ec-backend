package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/o-ga09/tutorial-ec-backend/app/application/product"
	"github.com/o-ga09/tutorial-ec-backend/pkg/validator"
)

type handler struct {
	saveProductUsecase *product.SaveProductUsecase
	fetchProductUsecase *product.FetchProductUseCase
}

func NewHandler(saveProductUsecas *product.SaveProductUsecase, fetchProductUsecase *product.FetchProductUseCase ) *handler {
	return &handler{
		saveProductUsecase: saveProductUsecas,
		fetchProductUsecase: fetchProductUsecase,
	}
}

// PostProducts godoc
// @Summary 商品を保存する
// @Tags products
// @Accept json
// @Produce json
// @Param request body []PostRequestParm ture "登録商品"
// @Success 201 {object} postProductResponse
// @Router /v1/products [post]
func(h handler) PostProducts(c *gin.Context) {
	var requestParam PostRequestParm

	err := c.ShouldBindJSON(requestParam)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"code": 403, "message": "Bad Request"})
		return
	}

	validate := validator.GetValidator()
	err = validate.Struct(requestParam)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"code": 403, "message": "Bad Request"})
		return
	}

	input := product.SaveProductUsecaseInputDto{
		OwnerID: requestParam.OwnerID,
		Name: requestParam.Name,
		Description: requestParam.Description,
		Price: requestParam.Price,
		Stock: requestParam.Stock,
	}

	dto, err := h.saveProductUsecase.Run(c,input)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"code": 500, "message": "Internal Server Eerror"})
		return

	}

	response := postProductResponse{
		ProductResponseModel{
			ID: dto.Id,
			OwnerID: dto.OwnerID,
			Name: dto.Name,
			Description: dto.Description,
			Price: dto.Price,
			Stock: dto.Stock,
		},
	}

	c.JSON(http.StatusOK,response)
}

func(h handler) GetProducts(c *gin.Context) {
	dtos, err := h.fetchProductUsecase.Run(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"code": 500, "message": "Internal Server Error"})
		return
	}

	var products []getProductResponse

	for _, dto := range dtos {
		products = append(products, getProductResponse{
			ProductResponseModel: &ProductResponseModel{
				ID: dto.ID,
				OwnerID: dto.OwnerID,
				Name: dto.Name,
				Description: dto.Description,
				Price: dto.Price,
				Stock: dto.Stock,
			},
			OwnerName: "owner name",
		})
	}

	c.JSON(http.StatusOK,products)
}