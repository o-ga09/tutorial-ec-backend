package product

import "context"

type FetchProductUseCase struct {
	productQueryService ProductQueryService
}

func NewFetchProductUseCase(productQueryService ProductQueryService) *FetchProductUseCase {
	return &FetchProductUseCase{productQueryService: productQueryService}
}

type FetchProductUseCaseDto struct {
	ID string
	OwnerID     string // 出品者ID
	Name        string // 商品名
	Description string // 商品の説明
	Price       int64  // 商品金額
	Stock       int    // 商品在庫
}

func(u FetchProductUseCase) Run(ctx context.Context) ([]*FetchProductUseCaseDto, error) {
	qsDtos, err := u.productQueryService.FetchProductList(ctx)
	if err != nil {
		return nil, err
	}

	var ucDtos []*FetchProductUseCaseDto

	for _, qsDto := range qsDtos {
		ucDtos = append(ucDtos,&FetchProductUseCaseDto{
			qsDto.ID,
			qsDto.OwnerID,
			qsDto.Name,
			qsDto.Description,
			qsDto.Price,
			qsDto.Stock,
		})
	}

	return ucDtos, nil
}