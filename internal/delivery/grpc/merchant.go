package grpcRoutes

import (
	repositoryRedis "github.com/techpartners-asia/payments-service/infrastructure/redis/repository"
	repositoryMerchant "github.com/techpartners-asia/payments-service/internal/modules/merchant/repository"
	usecaseMerchant "github.com/techpartners-asia/payments-service/internal/modules/merchant/usecase"
	merchantProto "github.com/techpartners-asia/payments-service/pkg/proto/merchant"
	"google.golang.org/grpc"
)

func RegisterMerchantServices(
	server grpc.ServiceRegistrar,
	merchantRepo repositoryMerchant.MerchantRepository,
	ebarimtRepo repositoryMerchant.EbarimtRepository,
	redisRepository repositoryRedis.RedisRepository,
) {
	merchantUsecase := usecaseMerchant.NewMerchantUsecase(merchantRepo, ebarimtRepo, redisRepository)
	merchantProto.RegisterMerchantServiceServer(server, merchantUsecase)
}
