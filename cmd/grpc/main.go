package grpcCmd

import (
	"log"
	"net"

	"github.com/techpartners-asia/payments-service/infrastructure/database"
	repositoryImpl "github.com/techpartners-asia/payments-service/infrastructure/database/repository"
	redisService "github.com/techpartners-asia/payments-service/infrastructure/redis"
	repositoryRedis "github.com/techpartners-asia/payments-service/infrastructure/redis/repository"
	grpcRoutes "github.com/techpartners-asia/payments-service/internal/delivery/grpc"
	"google.golang.org/grpc"
)

func Run() {

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()

	// * : Dependency Injection
	redisRepository := repositoryRedis.NewRedisRepository(redisService.Redis)
	paymentRepo := repositoryImpl.NewPaymentRepository(database.DB)
	merchantRepo := repositoryImpl.NewMerchantRepository(database.DB)
	ebarimtRepo := repositoryImpl.NewMerchantEbarimtRepository(database.DB)

	// * : Register Services
	grpcRoutes.RegisterServices(grpcServer, redisRepository, paymentRepo)
	grpcRoutes.RegisterMerchantServices(grpcServer, merchantRepo, ebarimtRepo, redisRepository)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
	log.Println("gRPC server is running on port 50051")
}
