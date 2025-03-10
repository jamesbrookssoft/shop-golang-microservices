package dtos

import (
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/utils"
)

type SearchProductsRequestDto struct {
	SearchText       string `query:"search" json:"search"`
	*utils.ListQuery `json:"listQuery"`
}
