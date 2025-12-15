package grpcRequestDTO

import (
	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"

	paymentProto "git.techpartners.asia/gateway-services/payment-service/pkg/proto/payment"
)

func ToEntity(req *paymentProto.PaymentCreateRequest) *entity.PaymentEntity {
	return &entity.PaymentEntity{
		Amount: float64(req.Amount),
		Status: entity.PaymentStatusPending,
	}
}
