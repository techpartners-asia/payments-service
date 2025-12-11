package grpcRoutes

import (
	repositoryRedis "github.com/techpartners-asia/payments-service/infrastructure/redis/repository"
	repositoryPayment "github.com/techpartners-asia/payments-service/internal/modules/payment/repository"
	usecasePayment "github.com/techpartners-asia/payments-service/internal/modules/payment/usecase"
	paymentProto "github.com/techpartners-asia/payments-service/pkg/proto/payment"
	"google.golang.org/grpc"
)

// PaymentHandler is your struct that implements the gRPC interface
// It would be defined in your internal/delivery/grpc package
type PaymentHandler struct {
	// ... your dependencies like use cases/modules
	paymentProto.UnimplementedPaymentServiceServer
}

// Compile-time check that PaymentHandler satisfies the generated interface.
var _ paymentProto.PaymentServiceServer = (*PaymentHandler)(nil)

// RegisterServices wires up all gRPC services to the provided server.
func RegisterServices(
	server grpc.ServiceRegistrar,
	redisRepository repositoryRedis.RedisRepository,
	paymentRepo repositoryPayment.PaymentRepository,
) {

	paymentUsecase := usecasePayment.NewPaymentUsecase(paymentRepo, redisRepository)
	paymentProto.RegisterPaymentServiceServer(server, paymentUsecase)
}
