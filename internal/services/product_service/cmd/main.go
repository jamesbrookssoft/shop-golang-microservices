package main

import (
	"github.com/go-playground/validator"
	gormpgsql "github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/gorm_pgsql"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/grpc"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/http"
	echoserver "github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/http/echo/server"
	httpclient "github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/http_client"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/logger"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/otel"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/rabbitmq"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/product_service/config"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/product_service/product/configurations"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/product_service/product/data/repositories"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/product_service/product/mappings"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/product_service/product/models"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/product_service/server"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				config.InitConfig,
				logger.InitLogger,
				http.NewContext,
				echoserver.NewEchoServer,
				grpc.NewGrpcClient,
				gormpgsql.NewGorm,
				otel.TracerProvider,
				httpclient.NewHttpClient,
				repositories.NewPostgresProductRepository,
				rabbitmq.NewRabbitMQConn,
				rabbitmq.NewPublisher,
				validator.New,
			),
			fx.Invoke(server.RunServers),
			fx.Invoke(configurations.ConfigMiddlewares),
			fx.Invoke(configurations.ConfigSwagger),
			fx.Invoke(func(gorm *gorm.DB) error {
				return gormpgsql.Migrate(gorm, &models.Product{})
			}),
			fx.Invoke(mappings.ConfigureMappings),
			fx.Invoke(configurations.ConfigEndpoints),
			fx.Invoke(configurations.ConfigProductsMediator),
		),
	).Run()
}
