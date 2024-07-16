package queries

import (
	"context"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/grpc"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/logger"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/rabbitmq"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/utils"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/product_service/product/data/contracts"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/product_service/product/dtos"
	dtosv1 "github.com/jamesbrookssoft/shop-golang-microservices/internal/services/product_service/product/features/searching_product/v1/dtos"
)

type SearchProductsHandler struct {
	log               logger.ILogger
	rabbitmqPublisher rabbitmq.IPublisher
	productRepository contracts.ProductRepository
	ctx               context.Context
	grpcClient        grpc.GrpcClient
}

func NewSearchProductsHandler(log logger.ILogger, rabbitmqPublisher rabbitmq.IPublisher,
	productRepository contracts.ProductRepository, ctx context.Context, grpcClient grpc.GrpcClient) *SearchProductsHandler {
	return &SearchProductsHandler{log: log, productRepository: productRepository, ctx: ctx, rabbitmqPublisher: rabbitmqPublisher, grpcClient: grpcClient}
}

func (c *SearchProductsHandler) Handle(ctx context.Context, query *SearchProducts) (*dtosv1.SearchProductsResponseDto, error) {

	products, err := c.productRepository.SearchProducts(ctx, query.SearchText, query.ListQuery)
	if err != nil {
		return nil, err
	}

	listResultDto, err := utils.ListResultToListResultDto[*dtos.ProductDto](products)
	if err != nil {
		return nil, err
	}

	return &dtosv1.SearchProductsResponseDto{Products: listResultDto}, nil
}
