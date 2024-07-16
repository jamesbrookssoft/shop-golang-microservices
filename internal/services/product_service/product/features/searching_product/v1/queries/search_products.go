package queries

import (
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/utils"
)

type SearchProducts struct {
	SearchText string `validate:"required"`
	*utils.ListQuery
}
