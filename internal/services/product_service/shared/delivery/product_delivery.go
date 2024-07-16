package delivery

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/grpc"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/logger"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/rabbitmq"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/product_service/config"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/product_service/product/data/contracts"
	"github.com/streadway/amqp"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type ProductDeliveryBase struct {
	Log               logger.ILogger
	Cfg               *config.Config
	RabbitmqPublisher rabbitmq.IPublisher
	ConnRabbitmq      *amqp.Connection
	HttpClient        *resty.Client
	JaegerTracer      trace.Tracer
	Gorm              *gorm.DB
	Echo              *echo.Echo
	GrpcClient        grpc.GrpcClient
	ProductRepository contracts.ProductRepository
	Ctx               context.Context
}
