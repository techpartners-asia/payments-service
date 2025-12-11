package grpcResponseDTO

import (
	"github.com/techpartners-asia/payments-service/infrastructure/database/entity"
	grpcMapDTO "github.com/techpartners-asia/payments-service/internal/modules/payment/dto/map/grpc"
	paymentProto "github.com/techpartners-asia/payments-service/pkg/proto/payment"
)

func ToCheckResponse(payment *entity.PaymentEntity) *paymentProto.PaymentCheckResponse {

	return &paymentProto.PaymentCheckResponse{
		Uid:    payment.UID,
		Amount: float32(payment.Amount),
		Status: grpcMapDTO.ToPaymentStatus(payment.Status),
	}
}
