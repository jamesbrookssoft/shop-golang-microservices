package dtos

import (
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/utils"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/product_service/product/dtos"
)

type GetProductsResponseDto struct {
	Products *utils.ListResult[*dtos.ProductDto]
}
