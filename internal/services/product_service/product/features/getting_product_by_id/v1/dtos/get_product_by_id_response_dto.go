package dtos

import (
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/product_service/product/dtos"
)

type GetProductByIdResponseDto struct {
	Product *dtos.ProductDto `json:"product"`
}
