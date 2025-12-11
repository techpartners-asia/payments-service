package grpcRequestDTO

import (
	"github.com/techpartners-asia/payments-service/infrastructure/database/entity"

	paymentProto "github.com/techpartners-asia/payments-service/pkg/proto/payment"
)

func ToEntity(req *paymentProto.PaymentCreateRequest) *entity.PaymentEntity {
	return &entity.PaymentEntity{
		Amount: float64(req.Amount),
		Status: entity.PaymentStatusPending,
	}
}
