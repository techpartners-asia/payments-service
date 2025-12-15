package grpcRoutes

import (
	repositoryRedis "git.techpartners.asia/gateway-services/payment-service/infrastructure/redis/repository"
	repositoryMerchant "git.techpartners.asia/gateway-services/payment-service/internal/modules/merchant/repository"
	usecaseMerchant "git.techpartners.asia/gateway-services/payment-service/internal/modules/merchant/usecase"
	merchantProto "git.techpartners.asia/gateway-services/payment-service/pkg/proto/merchant"
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
