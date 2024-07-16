package configurations

import (
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/grpc"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/pkg/logger"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/identity_service/config"
	identityservice "github.com/jamesbrookssoft/shop-golang-microservices/internal/services/identity_service/identity/grpc_server/protos"
	"github.com/jamesbrookssoft/shop-golang-microservices/internal/services/identity_service/identity/grpc_server/services"
)

func ConfigIdentityGrpcServer(cfg *config.Config, log logger.ILogger, grpcServer *grpc.GrpcServer) {

	identityGrpcService := services.NewIdentityGrpcServerService(cfg, log)

	identityservice.RegisterIdentityServiceServer(grpcServer.Grpc, identityGrpcService)
}
