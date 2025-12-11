package grpcResponseDTO

import (
	"github.com/techpartners-asia/payments-service/infrastructure/database/entity"
	grpcMapDTO "github.com/techpartners-asia/payments-service/internal/modules/payment/dto/map/grpc"
	paymentProto "github.com/techpartners-asia/payments-service/pkg/proto/payment"
)

// ToCreateResponse converts an internal Entity into a Protobuf response.
func ToCreateResponse(p *entity.PaymentEntity) *paymentProto.PaymentCreateResponse {
	return &paymentProto.PaymentCreateResponse{
		Uid:       p.UID,
		Amount:    float32(p.Amount),
		Status:    grpcMapDTO.ToPaymentStatus(p.Status),
		InvoiceID: p.RefInvoiceID,
	}
}
