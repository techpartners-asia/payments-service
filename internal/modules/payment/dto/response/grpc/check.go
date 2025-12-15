package grpcResponseDTO

import (
	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	grpcMapDTO "git.techpartners.asia/gateway-services/payment-service/internal/modules/payment/dto/map/grpc"
	paymentProto "git.techpartners.asia/gateway-services/payment-service/pkg/proto/payment"
)

func ToCheckResponse(payment *entity.PaymentEntity) *paymentProto.PaymentCheckResponse {

	return &paymentProto.PaymentCheckResponse{
		Uid:    payment.UID,
		Amount: float32(payment.Amount),
		Status: grpcMapDTO.ToPaymentStatus(payment.Status),
	}
}
