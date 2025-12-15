package grpcResponseDTO

import (
	"git.techpartners.asia/gateway-services/payment-service/infrastructure/database/entity"
	grpcMapDTO "git.techpartners.asia/gateway-services/payment-service/internal/modules/payment/dto/map/grpc"
	paymentProto "git.techpartners.asia/gateway-services/payment-service/pkg/proto/payment"
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
