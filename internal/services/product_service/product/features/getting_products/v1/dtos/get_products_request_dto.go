package dtos

import (
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/utils"
)

type GetProductsRequestDto struct {
	*utils.ListQuery
}
