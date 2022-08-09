package getting_product_by_id

import (
	"context"
	"fmt"
	"github.com/meysamhadeli/shop-golang-microservices/pkg/http_errors"
	"github.com/meysamhadeli/shop-golang-microservices/pkg/logger"
	"github.com/meysamhadeli/shop-golang-microservices/pkg/mapper"
	"github.com/meysamhadeli/shop-golang-microservices/services/products/config"
	"github.com/meysamhadeli/shop-golang-microservices/services/products/internal/products/contracts"
	"github.com/meysamhadeli/shop-golang-microservices/services/products/internal/products/dto"
	"github.com/meysamhadeli/shop-golang-microservices/services/products/internal/products/features/getting_product_by_id/dtos"
	"github.com/opentracing/opentracing-go"
)

type GetProductByIdHandler struct {
	log    logger.ILogger
	cfg    *config.Config
	pgRepo contracts.ProductRepository
}

func NewGetProductByIdHandler(log logger.ILogger, cfg *config.Config, pgRepo contracts.ProductRepository) *GetProductByIdHandler {
	return &GetProductByIdHandler{log: log, cfg: cfg, pgRepo: pgRepo}
}

func (q *GetProductByIdHandler) Handle(ctx context.Context, query *GetProductById) (*dtos.GetProductByIdResponseDto, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "getProductByIdHandler.Handle")
	defer span.Finish()

	product, err := q.pgRepo.GetProductById(ctx, query.ProductID)

	if err != nil {
		return nil, http_errors.NewNotFoundError(fmt.Sprintf("product with id %s not found", query.ProductID))
	}

	productDto, err := mapper.Map[*dto.ProductDto](product)
	if err != nil {
		return nil, err
	}

	return &dtos.GetProductByIdResponseDto{Product: productDto}, nil
}
