package main

import (
	"github.com/go-playground/validator"
	gormpgsql "github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/gorm_pgsql"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/grpc"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/http"
	echoserver "github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/http/echo/server"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/http_client"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/logger"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/oauth2"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/otel"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/rabbitmq"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/identity_service/config"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/identity_service/identity/configurations"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/identity_service/identity/data/repositories"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/identity_service/identity/data/seeds"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/identity_service/identity/mappings"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/identity_service/identity/models"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/identity_service/server"
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
				grpc.NewGrpcServer,
				gormpgsql.NewGorm,
				otel.TracerProvider,
				httpclient.NewHttpClient,
				repositories.NewPostgresUserRepository,
				rabbitmq.NewRabbitMQConn,
				rabbitmq.NewPublisher,
				validator.New,
			),
			fx.Invoke(server.RunServers),
			fx.Invoke(configurations.ConfigMiddlewares),
			fx.Invoke(configurations.ConfigSwagger),
			fx.Invoke(func(gorm *gorm.DB) error {
				err := gormpgsql.Migrate(gorm, &models.User{})
				if err != nil {
					return err
				}
				return seeds.DataSeeder(gorm)
			}),
			fx.Invoke(mappings.ConfigureMappings),
			fx.Invoke(configurations.ConfigEndpoints),
			fx.Invoke(configurations.ConfigUsersMediator),
			fx.Invoke(oauth2.RunOauthServer),
		),
	).Run()
}
