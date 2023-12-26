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
	ownerID     string // 出品者ID
	name        string // 商品名
	description string // 商品の説明
	price       int64  // 商品金額
	stock       int    // 商品在庫
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
			qsDto.ownerID,
			qsDto.name,
			qsDto.description,
			qsDto.price,
			qsDto.stock,
		})
	}

	return ucDtos, nil
}